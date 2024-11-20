package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

// GlobalException exception handling
func GlobalException(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.IsAborted() {
				return
			}

			var stacktrace string
			for i := 1; ; i++ {
				_, f, l, got := runtime.Caller(i)
				if !got {
					break
				}
				stacktrace += fmt.Sprintf("%s:%d\n", f, l)
			}

			errMsg := fmt.Sprintf("Exception message: %s", err)
			// when stack finishes
			logMessage := fmt.Sprintf("Recovery from exception：('%s')\n", ctx.HandlerName())
			logMessage += errMsg + "\n"
			logMessage += fmt.Sprintf("\n%s", stacktrace)
			// 打印错误日志
			//logs.GetLogger().Error(logMessage)
			//// 返回错误信息
			ctx.JSON(500, logMessage)
			ctx.Status(500)
			ctx.Abort()
		}
	}()
	ctx.Next()
}
