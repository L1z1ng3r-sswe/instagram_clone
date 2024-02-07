package middleware

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/service"
	"github.com/L1z1ng3r-sswe/instagram_clone/app/pkg/logging"
	"github.com/gin-gonic/gin"
)

func IsAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		access_token := ctx.GetHeader("Authorization")
		logger := logging.GetLogger()

		_, _, line, _ := runtime.Caller(0)

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Access-token not found"})
			logger.Err("Unathorized", "Token not found", "app/internal/service/"+"is_auth_mw"+" line: "+strconv.Itoa(line))
			return
		}

		err, errKey, errMsg, code, fileInfo, userId := service.IsTokenValid(access_token, false)
		if err != nil {
			ctx.AbortWithStatusJSON(code, gin.H{errKey: errMsg})
			logger.Err(errKey, errMsg, fileInfo)
			return
		}

		fmt.Println("from mw: id: ", userId)

		ctx.Set("user_id", userId)

		ctx.Next()
	}
}
