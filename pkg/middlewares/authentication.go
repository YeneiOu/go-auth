package middlewares

import (
	"clean-arc/configs"
	"clean-arc/helpers"
	"clean-arc/modules/entities"
	"clean-arc/pkg/utils"
	"context"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func JwtAuthentication(cfg *configs.Configs) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.WithValue(c.Context(), entities.MiddlewaresCon, time.Now().UnixMilli())
		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		log.Printf("called:\t%v", utils.Trace())
		defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.MiddlewaresCon).(int64)))
		if accessToken == "" {
			log.Println("error: authorization is empty")
			return helpers.RespondWithJSON(c, fiber.ErrUnauthorized.Message, fiber.ErrUnauthorized.Code, "authorization is empty", nil)
		}

		userId, err := utils.JwtExtractPayload(ctx, cfg, "user_id", accessToken)

		if err != nil {
			return helpers.RespondWithJSON(c, fiber.ErrUnauthorized.Message, fiber.ErrUnauthorized.Code, err.Error(), nil)
		}

		c.Locals("user_id", userId)
		username, err := utils.JwtExtractPayload(ctx, cfg, "username", accessToken)

		if err != nil {
			return helpers.RespondWithJSON(c, fiber.ErrUnauthorized.Message, fiber.ErrUnauthorized.Code, err.Error(), nil)
		}
		c.Locals("username", username)

		userRole, err := utils.JwtExtractPayload(ctx, cfg, "role", accessToken)

		if err != nil {
			return helpers.RespondWithJSON(c, fiber.ErrUnauthorized.Message, fiber.ErrUnauthorized.Code, err.Error(), nil)
		}
		//Set role into Fiber cache in case to use with another function from next
		log.Printf("userRole extracted from JWT: %s", userRole)
		c.Locals("role", userRole)

		paramsUserId := c.Params("user_id")
		if paramsUserId != "" && paramsUserId != userId {
			return helpers.RespondWithJSON(c, fiber.ErrUnauthorized.Message, fiber.ErrUnauthorized.Code, "error, have no permission to access", nil)
		}

		return c.Next()
	}

}
