package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jrione/gin-crud/domain"
)

func CreateAccessToken(user *domain.User, secret string, expire int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expire)).Unix()
	claims := &domain.JWTClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return accessToken, err

}

func IsAuthorized(authToken string, secret string) (bool, error) {
	_, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Cannot signing method with algorithm: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Errorf("Error parsing token: %v", err)
		return false, err
	}
	return true, nil

}

func GetUsernameFromToken(authToken, secret string) (string, error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Cannot signing method with algorithm: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("Error Occured: %v")
	}
	return claim["username"].(string), nil
}
