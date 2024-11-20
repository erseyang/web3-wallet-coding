package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisClient *redis.Pool

func RedisConnect(host, port, password string, db int) {
	redisClient = &redis.Pool{
		MaxIdle:         100,
		MaxActive:       2014,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				return nil, err
			}
			_, err = c.Do("AUTH", password)
			if err != nil {
				return nil, err
			}
			_, err = c.Do("SELECT", db)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func GetRedisPool() *redis.Pool {
	return redisClient
}
