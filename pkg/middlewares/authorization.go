package middlewares

import (
	"clean-arc/helpers"
	"clean-arc/modules/entities"
	"clean-arc/pkg/utils"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Authorization(roles ...entities.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.WithValue(c.Context(), entities.MiddlewaresCon, time.Now().UnixMilli())
		log.Printf("called:\t%v", utils.Trace())
		defer log.Printf("return:\t%v time:%v ms", utils.Trace(), utils.CallTimer(ctx.Value(entities.MiddlewaresCon).(int64)))

		rolesMap := map[entities.UserRole]int{
			"user":    1, // 2^0
			"manager": 2, // 2^1
			"admin":   4, // 2^2
		}
		// Find a summation of expected role
		var sum int
		for i := range roles {
			sum += rolesMap[roles[i]]
		}

		roleBinary := utils.BinaryConvertor(sum, len(rolesMap))

		roleInterface := c.Locals("role")
		if roleInterface == nil {

			return helpers.RespondWithJSON(c, fiber.ErrUnauthorized.Message, fiber.StatusUnauthorized, "Error, role not set in context", nil)
		}
		userRole := rolesMap[entities.UserRole(roleInterface.(string))]

		userRoleBinary := utils.BinaryConvertor(userRole, len(rolesMap))

		// Bitwise operator to compare a role which can access or not
		for i := 0; i < len(rolesMap); i++ {
			if roleBinary[i]&userRoleBinary[i] == 1 {
				return c.Next()
			}
		}
		return helpers.RespondWithJSON(c, fiber.ErrUnauthorized.Message, fiber.StatusUnauthorized, "Error, have no permission to access", nil)
	}
}
