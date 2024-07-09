package helpers

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		token, err = getTokenWithPrivateKey(context)
		if err != nil {
			return err
		}
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
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenWithPrivateKey(context *gin.Context) (*jwt.Token, error) {
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

func GenerateJWT(userId int, fullname string, authToken string) (string, error) {
	privateKey := []byte(os.Getenv("JWT_PRIVATE_KEY"))
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         userId,
		"fullname":   fullname,
		"auth_token": authToken,
		"exp":        time.Now().Add(time.Hour * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)
}

func ParseToken(jwtToken string) (*jwt.Token, error) {
	privateKey := []byte(os.Getenv("JWT_PRIVATE_KEY"))
	splitToken := strings.Split(jwtToken, " ")
	token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}
