package router

import (
	"com.wallet/coding/config"
	"com.wallet/coding/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	if config.Instance.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	userController := controller.NewUserController()
	r.GET("/user/login", userController.Login)
	r.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{"Message": "Hello World"})
	})
	return r
}
