package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/sso-jrione/config"
	"github.com/jrione/sso-jrione/controller"
	_userRepo "github.com/jrione/sso-jrione/repository/postgres"
	_userUseCase "github.com/jrione/sso-jrione/usecase"
)

func NewUserRoute(env *config.Config, timeout time.Duration, dbclient *sql.DB, gr *gin.RouterGroup) {
	userRepo := _userRepo.NewUserRepository(dbclient)
	userUseCase := _userUseCase.NewUserUseCase(userRepo, timeout)

	userController := controller.UserController{
		UserUseCase: userUseCase,
	}

	gr.GET("/user", userController.GetAllUser)

}
