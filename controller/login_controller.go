package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/config"
	"github.com/jrione/gin-crud/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUseCase domain.LoginUseCase
	Env          *config.Config
}

func checkHashPass(hashed string, realPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(realPass))
	return err == nil
}

func (l LoginController) Login(gctx *gin.Context) {
	req := domain.LoginRequest{}
	err := gctx.ShouldBindJSON(&req)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Username and Password Required!",
		})
		return
	}

	data, err := l.LoginUseCase.CheckUser(gctx, req.Username)
	if err != nil {
		gctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusUnauthorized,
			"err":  "Incorrect Username Or Password",
		})
		return
	}
	if req.Username != data.Username && !checkHashPass(data.Password, req.Password) {
		gctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"err":  "Incorrect Username Or Password",
		})
		return
	}

	accessToken, err := l.LoginUseCase.CreateAccessToken(data, l.Env.Server.AccessTokenSecret, l.Env.Server.AccessTokenExpiry)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"cause": err,
		})
		return
	}
	refreshToken, err := l.LoginUseCase.CreateAccessToken(data, l.Env.Server.RefreshTokenSecret, l.Env.Server.RefreshTokenExpiry)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"cause": err,
		})
		return
	}

	resp := &domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	gctx.JSON(http.StatusOK, resp)
}

func (l LoginController) GetLogin(gctx *gin.Context) {

	data, err := l.LoginUseCase.CheckUser(gctx, "njir")
	if err != nil {
		gctx.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusUnauthorized,
			"err":  "Incorrect Username Or Password",
		})
		return
	}

	accessToken, err := l.LoginUseCase.CreateAccessToken(data, l.Env.Server.AccessTokenSecret, l.Env.Server.AccessTokenExpiry)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"cause": err,
		})
		return
	}

	refreshToken, err := l.LoginUseCase.CreateAccessToken(data, l.Env.Server.AccessTokenSecret, l.Env.Server.AccessTokenExpiry)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"cause": err,
		})
		return
	}

	resp := &domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	gctx.JSON(http.StatusOK, resp)
}
