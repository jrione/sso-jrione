package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/sso-jrione/config"
	"github.com/jrione/sso-jrione/controller"
	userRepo "github.com/jrione/sso-jrione/repository/postgres"
	userUseCase "github.com/jrione/sso-jrione/usecase"
)

func NewRegisterRoute(env *config.Config, timeout time.Duration, dbclient *sql.DB, gr *gin.RouterGroup) {
	ur := userRepo.NewUserRepository(dbclient)
	uu := userUseCase.NewUserUseCase(ur, timeout)

	uc := &controller.RegisterController{
		UserUseCase: uu,
	}

	gr.POST("/register", uc.Register)

}
