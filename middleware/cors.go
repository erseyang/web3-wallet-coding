package middleware

import "github.com/gin-gonic/gin"

func Cors(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request.Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "SessionId, Origin, X-Requested-With, Content-Type, Accept, Authorization, X-Request-Id, Img-Code, Imgcode-No, Mobile, Smscode")
		ctx.Status(204)
		return
	}
	ctx.Next()
}
