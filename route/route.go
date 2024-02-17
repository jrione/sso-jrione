package route

import (
	"database/sql"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/config"
	middleware "github.com/jrione/gin-crud/middleware"
)

func SetupRoute(env *config.Config, timeout time.Duration, dbclient *sql.DB, r *gin.Engine) {
	publicRouter := r.Group("")
	NewTestRoute(env, timeout, publicRouter)

	authRouter := r.Group("/auth")
	store := cookie.NewStore([]byte("secret"))

	// authRouter.Use(middleware.AuthMiddleware(env))
	authRouter.Use(sessions.Sessions("sess", store))
	NewLoginRoute(env, timeout, dbclient, authRouter)

	authRouter.Use(middleware.SessionMiddleware())
	NewUserRoute(env, timeout, dbclient, authRouter)

}
