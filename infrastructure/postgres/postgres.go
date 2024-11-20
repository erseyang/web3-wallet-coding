package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
)

var (
	db     *sql.DB
	dbLock sync.Mutex
)

func Instance(host, user, pass, name string, port int) *sql.DB {
	if db != nil {
		return db
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	if db == nil {
		driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&interpolateParams=true",
			user, pass, host, port, name)
		db, _ = open("postgres", driverSource)
	}
	return db
}

func open(drverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(drverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)
	return db, nil
}
