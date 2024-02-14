package auth

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-Key")
		if apiKey == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "Unauthorized",
			})
			return
		}
	}
}
