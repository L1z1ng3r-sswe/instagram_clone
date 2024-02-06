package logging

import (
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()

		logger := GetLogger()

		ctx.Set("logger", logger)

		logger.Request(ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), time.Since(start))
	}
}
