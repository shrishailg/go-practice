package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-practice/config"
)

// request Authentivation Check
func Authentication(ctx *gin.Context) {
	header := ctx.Request.Header

	if header.Get(config.Config.GetString("authentication.header_key")) != config.Config.GetString("authentication.header_value") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid authentication",
		})
	}

	ctx.Next()
}

func AuthenticationV2(ctx *gin.Context) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		header := ctx.Request.Header

		if header.Get("token") != "app_user" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authentication",
			})

			return
		}

		ctx.Next()
	}
}

// add header to response
func SetResponseHeader(ctx *gin.Context) {
	header := ctx.Writer.Header()

	header.Set("wtf", "executor")
}
