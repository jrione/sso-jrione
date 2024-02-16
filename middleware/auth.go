package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/config"
)

func AuthMiddleware(env *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-Key")
		if apiKey != env.Server.XApiKey {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "Unauthorized",
			})
			return
		}
	}
}
