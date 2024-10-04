package controllers

import (
	"clean-arc/configs"
	"clean-arc/helpers"
	"clean-arc/pkg/middlewares"

	"clean-arc/modules/entities"

	"github.com/gofiber/fiber/v2"
)

type authCon struct {
	Cfg     *configs.Configs
	AuthUse entities.AuthUsecase
}

func NewAuthController(r fiber.Router, cfg *configs.Configs, authUse entities.AuthUsecase) {
	handleController := &authCon{
		Cfg:     cfg,
		AuthUse: authUse,
	}
	r.Post("/login", handleController.Login)
	r.Get("/auth-test", middlewares.JwtAuthentication(handleController.Cfg), handleController.AuthTest)
	r.Get("/authentication", middlewares.JwtAuthentication(handleController.Cfg), handleController.Authentication)
	r.Get("/authorization", middlewares.JwtAuthentication(handleController.Cfg), middlewares.Authorization("user"), handleController.Authorization)
}

func (h *authCon) Login(c *fiber.Ctx) error {
	req := new(entities.UsersCredentials)
	if err := c.BodyParser(req); err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, err.Error(), nil)
	}

	res, err := h.AuthUse.Login(h.Cfg, req)
	if err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, err.Error(), nil)
	}

	return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "", res)
}
func (h *authCon) AuthTest(c *fiber.Ctx) error {
	id := c.Locals("user_id")
	username := c.Locals("username")
	return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "", map[string]interface{}{
		"id":       id,
		"username": username,
	})
}
func (h *authCon) Authentication(c *fiber.Ctx) error {
	return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "", nil)
}

func (h *authCon) Authorization(c *fiber.Ctx) error {
	return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "", nil)
}
