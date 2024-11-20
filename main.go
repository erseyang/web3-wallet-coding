package main

import (
	"com.wallet/coding/config"
	"com.wallet/coding/infrastructure/logs"
	"com.wallet/coding/router"
	"flag"
	"fmt"
	"net/http"
	"time"
)

var configFile = flag.String("config", "./config/wallet.yaml", "Configure file path")

func init() {

	flag.Parse()

	config.Init(*configFile)

	fmt.Println(config.Instance.Port)

}

func main() {
	r := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", 8078),
		Handler:        r,
		ReadTimeout:    time.Duration(60) * time.Second,
		WriteTimeout:   time.Duration(60) * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	err := s.ListenAndServe()
	if err != nil {
		logs.GetLogger().Error(err.Error())
	}
}
