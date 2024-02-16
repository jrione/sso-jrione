package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/domain"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (u UserController) GetAllUser(gctx *gin.Context) {
	listUser, err := u.UserUseCase.GetAll(gctx)
	if err != nil {
		log.Fatal(err)
	}
	gctx.JSON(http.StatusOK, listUser)
}
