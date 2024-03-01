package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/config"
	"github.com/jrione/gin-crud/controller"

	repo "github.com/jrione/gin-crud/repository/postgres"
	useCase "github.com/jrione/gin-crud/usecase"
)

func NewLoginRoute(env *config.Config, timeout time.Duration, dbclient *sql.DB, gr *gin.RouterGroup) {
	lr := repo.NewUserRepository(dbclient)
	lu := useCase.NewLoginUseCase(lr, timeout)

	rr := repo.NewRefreshTokenRepository(dbclient)
	ru := useCase.NewRefreshTokenUseCase(rr, timeout)

	lc := &controller.LoginController{
		LoginUseCase:        lu,
		RefreshTokenUseCase: ru,
		Env:                 env,
	}

	gr.POST("/login", lc.Login)
	gr.POST("/refresh", lc.RefreshToken)
	gr.GET("/getLogin", lc.GetLogin)
}
