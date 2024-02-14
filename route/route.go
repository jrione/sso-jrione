package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/config"
	middleware "github.com/jrione/gin-crud/middleware"
)

func SetupRoute(env *config.Config, timeout time.Duration, dbclient *sql.DB, r *gin.Engine) {
	publicRouter := r.Group("")

	NewTestRoute(env, timeout, publicRouter)

	privateRouter := r.Group("")
	privateRouter.Use(middleware.AuthMiddleware())
	NewStudentRoute(env, timeout, dbclient, privateRouter)
}
