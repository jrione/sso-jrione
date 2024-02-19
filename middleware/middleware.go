package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/config"
)

func AuthMiddleware(env *config.Config) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		apiKey := gctx.GetHeader("X-API-Key")
		if apiKey != env.Server.XApiKey {
			gctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token X-Api-Key not provided",
			})
			return
		}
	}
}

func SessionMiddleware() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		sess := sessions.Default(gctx)
		isLoggedIn := sess.Get("IsLoggedIn")
		fmt.Println("sess", isLoggedIn)
		if isLoggedIn != true {
			gctx.Redirect(http.StatusTemporaryRedirect, "/auth/login")
			gctx.Abort()
		} else {
			gctx.Next()
		}
	}
}
