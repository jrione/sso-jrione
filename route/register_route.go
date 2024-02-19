package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/config"
	"github.com/jrione/gin-crud/controller"
	userRepo "github.com/jrione/gin-crud/repository/postgres"
	userUseCase "github.com/jrione/gin-crud/usecase"
)

func NewRegisterRoute(env *config.Config, timeout time.Duration, dbclient *sql.DB, gr *gin.RouterGroup) {
	ur := userRepo.NewUserRepository(dbclient)
	uu := userUseCase.NewUserUseCase(ur, timeout)

	uc := &controller.RegisterController{
		UserUseCase: uu,
	}

	gr.POST("/register", uc.Register)

}
