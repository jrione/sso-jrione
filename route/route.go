package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/sso-jrione/config"
	middleware "github.com/jrione/sso-jrione/middleware"
)

func SetupRoute(env *config.Config, timeout time.Duration, dbclient *sql.DB, r *gin.Engine) {
	publicRouter := r.Group("/")
	NewTestRoute(env, timeout, publicRouter)

	authRouter := r.Group("/auth")
	NewRegisterRoute(env, timeout, dbclient, authRouter)
	NewLoginRoute(env, timeout, dbclient, authRouter)

	sessionRouter := r.Group("")
	sessionRouter.Use(middleware.JWTMiddleware(env.Server.AccessTokenSecret))
	NewUserRoute(env, timeout, dbclient, sessionRouter)

}
