package controllers

import (
	"clean-arc/helpers"
	"clean-arc/modules/entities"

	"github.com/gofiber/fiber/v2"
)

type usersController struct {
	UsersUse entities.UsersUsecase
}

func NewUsersController(r fiber.Router, usersUse entities.UsersUsecase) {
	controllers := &usersController{
		UsersUse: usersUse,
	}
	r.Post("/register", controllers.Register)
}

func (h *usersController) Register(c *fiber.Ctx) error {
	req := new(entities.UsersRegisterReq)
	if err := c.BodyParser(req); err != nil {

		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, err.Error(), nil)
	}

	res, err := h.UsersUse.Register(req)
	if err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, err.Error(), nil)
	}

	return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "Create Success", res)
}
