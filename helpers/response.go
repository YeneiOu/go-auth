package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RespondWithJSON(c *fiber.Ctx, status string, statusCode int, message string, data interface{}) error {
	response := fiber.Map{
		"status":  status,
		"code":    statusCode,
		"message": message,
		"data":    data,
	}
	fmt.Printf("response: %v\n", response)
	return c.Status(statusCode).JSON(response)
}
