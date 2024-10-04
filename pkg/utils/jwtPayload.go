package utils

import (
	"clean-arc/configs"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func JwtExtractPayload(ctx context.Context, cfg *configs.Configs, fieldName string, tokenString string) (any, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error, unexpected signing method: %v", token.Header["alg"])
		}
		fmt.Println("1234")
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims[fieldName], nil
	}
	return "", errors.New("error, token is invalid")
}
