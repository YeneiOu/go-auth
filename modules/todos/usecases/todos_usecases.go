package usecases

import (
	"clean-arc/modules/entities"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type todosUsecase struct {
	TodosRepo entities.TodosRepository
}

// Constructor
func NewTodosUsecase(todoRepo entities.TodosRepository) entities.TodosUsecase {
	return &todosUsecase{
		TodosRepo: todoRepo,
	}
}

func (r *todosUsecase) GetAllTodos(c *fiber.Ctx, req *entities.TodosReq) ([]entities.TodosRes, error) {
	allTodo, err := r.TodosRepo.GetAllTodos(req)
	if err != nil {
		return nil, err
	}
	return allTodo, nil
}

func (r *todosUsecase) AddTodos(req *entities.TodosReq) (*entities.TodosRes, error) {

	inputNewTodo := entities.TodosReq{
		Title:    req.Title,
		Complete: false,
	}
	fmt.Println("inputNewTodo", inputNewTodo)
	newTodo, err := r.TodosRepo.AddTodos(&inputNewTodo)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}

func (r *todosUsecase) UpdateAllTodos(req []entities.TodosReq) ([]entities.TodosRes, error) {

	if len(req) == 0 {
		return nil, nil
	}
	newTodos, err := r.TodosRepo.UpdateAllTodos(req)
	if err != nil {
		return nil, nil
	}
	return newTodos, nil
}

func (r *todosUsecase) DeleteTodos(id int64) error {
	err := r.TodosRepo.DeleteTodos(id)
	if err != nil {
		return err
	}
	return nil
}
