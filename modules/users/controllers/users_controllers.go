package controllers

import (
	"clean-arc/configs"
	"clean-arc/helpers"
	"clean-arc/modules/entities"
	"clean-arc/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

type usersController struct {
	Cfg      *configs.Configs
	UsersUse entities.UsersUsecase
}

func NewUsersController(r fiber.Router, usersUse entities.UsersUsecase) {
	controllers := &usersController{
		UsersUse: usersUse,
	}
	r.Post("/", middlewares.JwtAuthentication(controllers.Cfg), middlewares.Authorization("user"), controllers.GetAllUsers)
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
func (h *usersController) GetAllUsers(c *fiber.Ctx) error {

	req := new(entities.GetAllUserReq)
	res, err := h.UsersUse.GetAllUsers(c, req)

	if err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, err.Error(), nil)
	}
	if len(res) == 0 {
		return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "No users found", []string{})
	}
	return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "Get All User", res)
}
