package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/config"
	"github.com/jrione/gin-crud/controller"

	loginRepo "github.com/jrione/gin-crud/repository/postgres"
	loginUseCase "github.com/jrione/gin-crud/usecase"
)

func NewLoginRoute(env *config.Config, timeout time.Duration, dbclient *sql.DB, gr *gin.RouterGroup) {
	lr := loginRepo.NewUserRepository(dbclient)
	lu := loginUseCase.NewLoginUseCase(lr, timeout)

	lc := &controller.LoginController{
		LoginUseCase: lu,
		Env:          env,
	}

	gr.POST("/login", lc.Login)
	gr.GET("/getLogin", lc.GetLogin)
}
