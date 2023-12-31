package middlewares

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	if !ok || !token.Valid {
		return err
	}

	return errors.New("invalid token provided")
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	privateKey := []byte(os.Getenv("JWT_PRIVATE_KEY"))
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func getTokenFromRequestSample(context *gin.Context) (string, error) {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1], nil
	}
	return "", errors.New("invalid token provided")
}

func CurrentDriver(context *gin.Context) error {
	err := ValidateJWT(context)
	if err != nil {
		return err
	}
	// token, _ := getToken(context)
	// claims, _ := token.Claims.(jwt.MapClaims)
	// userId := claims["id"].(string)

	// user, err := model.UserMetaFindById(userId)
	// return user, err
	return nil
}
