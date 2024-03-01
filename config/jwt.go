package config

import (
	"fmt"
	"log"
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

func CreateRefreshToken(user *domain.User, secret string, expire int) (refreshToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expire)).Unix()
	claims := &domain.JWTClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", nil
	}
	return refreshToken, nil
}

func IsAuthorized(authToken string, secret string) (bool, error) {
	_, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Cannot signing method with algorithm: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		if validErr, ok := err.(*jwt.ValidationError); ok {
			if validErr.Errors&jwt.ValidationErrorMalformed != 0 {
				log.Printf("ErrTokenMalformed: %v", err)
				return false, err
			} else if validErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				log.Printf("ErrTokenExpired: %v", err)
				return false, err
			}
		}
		log.Printf("Error parsing token: %v", err)
		return false, err
	}
	return true, nil

}

func GetUsernameFromClaim(authToken string, secret string) (user string, err error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Cannot signing method with algorithm: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		if validErr, ok := err.(*jwt.ValidationError); ok {
			if validErr.Errors&jwt.ValidationErrorMalformed != 0 {
				log.Printf("ErrTokenMalformed: %v", err)
				return "", err
			} else if validErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				log.Printf("ErrTokenExpired: %v", err)
				return "", err
			}
		}
		log.Printf("Error parsing token: %v", err)
		return "false", err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["username"].(string), nil
}

func IsExpired(authToken, secret string) (string, string, bool) {
	return "", "", false
}
