package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "dirahasiakanyaom"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Printf("Error generating token : %v", err.Error())
		return err.Error()
	}

	return signedToken
}

func VerifyToken(ctx *gin.Context) (jwt.MapClaims, error) {
	errResponse := errors.New("Sign in to proceed")
	headerToken := ctx.GetHeader("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	var mapClaims jwt.MapClaims

	if v, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, errResponse
	} else {
		mapClaims = v
	}

	return mapClaims, nil
}
