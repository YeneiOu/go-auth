package utils

import "github.com/gofiber/fiber/v2"

type UsernameBinding struct {
	Username string `json:"username" form:"username" binding:"required"`
}

func BindingUsername(c *fiber.Ctx) (string, error) {
	var obj UsernameBinding
	if err := c.BodyParser(&obj); err != nil {
		return obj.Username, err
	}

	return obj.Username, nil
}
