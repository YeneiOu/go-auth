package controllers

import (
	"clean-arc/configs"
	"clean-arc/helpers"
	"clean-arc/modules/entities"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type todosController struct {
	Cfg          *configs.Configs
	TodosUsecase entities.TodosUsecase
}

func NewTodosController(r fiber.Router, todosUse entities.TodosUsecase) {
	controllers := &todosController{
		TodosUsecase: todosUse,
	}
	r.Get("/", controllers.GetAllTodos)
	r.Post("/add-todo", controllers.AddTodos)
	r.Patch("/update-todo", controllers.UpdateAllTodos)
	r.Delete("/delete-todo", controllers.DeleteTodos)
}

func (h *todosController) GetAllTodos(c *fiber.Ctx) error {

	req := new(entities.TodosReq)
	res, err := h.TodosUsecase.GetAllTodos(c, req)

	if err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, err.Error(), nil)
	}
	if len(res) == 0 {
		return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "No todo items", []string{})
	}
	return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "Get All Todo", res)
}
func (h *todosController) AddTodos(c *fiber.Ctx) error {

	req := new(entities.TodosReq)
	if err := c.BodyParser(req); err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, "Invalid request payload", nil)
	}

	res, err := h.TodosUsecase.AddTodos(req)
	if err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, err.Error(), nil)
	}

	return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "Todo Added Successfully", res)

}
func (h *todosController) UpdateAllTodos(c *fiber.Ctx) error {

	// Create a pointer to a slice of TodosReq
	req := new([]entities.TodosReq)
	fmt.Println("req", req)
	// Parse the body into the request slice
	if err := c.BodyParser(req); err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, "Invalid request payload", nil)
	}

	// Dereference the pointer before passing it to the usecase
	res, err := h.TodosUsecase.UpdateAllTodos(*req)
	if err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, err.Error(), nil)
	}

	// Return a success response
	return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "Todos updated successfully", res)
}
func (h *todosController) DeleteTodos(c *fiber.Ctx) error {

	req := new(entities.TodosReq)
	fmt.Println("req", req)

	if err := c.BodyParser(req); err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, "Invalid request payload", nil)
	}

	err := h.TodosUsecase.DeleteTodos(req.ID)
	if err != nil {
		return helpers.RespondWithJSON(c, fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code, err.Error(), nil)
	}

	// Return a success response
	return helpers.RespondWithJSON(c, "OK", fiber.StatusOK, "Todos Delete successfully", nil)
}
