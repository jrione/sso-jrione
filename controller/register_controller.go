package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/domain"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {
	UserUseCase domain.UserUseCase
}

func hashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (reg RegisterController) Register(gctx *gin.Context) {
	req := &domain.UserRequest{}
	err := gctx.ShouldBindJSON(&req)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Payload Required!",
		})
		return
	}
	req.Password, err = hashPassword(req.Password)
	if err != nil {
		log.Fatalf("Error: Cannot hash password. Cause: %v", err)
	}

	gctx.JSON(http.StatusOK, gin.H{
		"data": req,
	})
}
