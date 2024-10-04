package helpers

import (
	"clean-arc/modules/entities"
	"github.com/gofiber/fiber/v2"
)

func RespondWithJSON(c *fiber.Ctx, status string, statusCode int, message string, data interface{}) error {
	response := entities.Responses{
		Status:  status,
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
	return c.Status(statusCode).JSON(response)
}
