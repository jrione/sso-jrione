package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/domain"
)

type LoginController struct {
	LoginUseCase domain.LoginUseCase
}

func (l LoginController) Login(gctx *gin.Context) {
	req := domain.LoginRequest{}
	err := gctx.BindJSON(&req)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "400 Bad Request",
			"error": err.Error(),
		})
		return
	}
	data, err := l.LoginUseCase.CheckUser(gctx, "njir")
	if err != nil {
		log.Fatal(err)
	}
	gctx.JSON(http.StatusOK, data)
}
