package middlewares

import (
	"net/http"

	"github.com/FarzadMohtasham/EventV8/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Not Authorized",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Not Authorized",
		})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
