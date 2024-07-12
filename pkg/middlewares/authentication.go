package middlewares

import (
	"clean-arc/helpers"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if accessToken == "" {
			log.Println("error: authorization is empty")
			return helpers.RespondWithJSON(c, fiber.ErrUnauthorized.Message, fiber.ErrUnauthorized.Code, "authorization is empty", nil)
		}
		secretKey := os.Getenv("JWT_SECRET_KEY")
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error, unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})
		if err != nil {
			log.Panicln(err.Error())
			return helpers.RespondWithJSON(c, fiber.ErrUnauthorized.Message, fiber.ErrUnauthorized.Code, "error, unauthorized", nil)
		}
		if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("user", claim)
			c.Locals("user_id", claim["user_id"])
			c.Locals("username", claim["username"])
			return c.Next()
		}
		return helpers.RespondWithJSON(c, fiber.ErrUnauthorized.Message, fiber.ErrUnauthorized.Code, "error, unauthorized", nil)
	}

}
