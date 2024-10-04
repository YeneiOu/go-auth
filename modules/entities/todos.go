package entities

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type TodosUsecase interface {
	GetAllTodos(c *fiber.Ctx, req *TodosReq) ([]TodosRes, error)
	AddTodos(req *TodosReq) (*TodosRes, error)
	UpdateAllTodos(req []TodosReq) ([]TodosRes, error)
	DeleteTodos(id int64) error
}

type TodosRepository interface {
	GetAllTodos(req *TodosReq) ([]TodosRes, error)
	AddTodos(req *TodosReq) (*TodosRes, error)
	UpdateAllTodos(req []TodosReq) ([]TodosRes, error)
	DeleteTodos(id int64) error
}

type TodosReq struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Title     string    `json:"title" db:"title"`
	Complete  bool      `json:"complete" db:"complete"`
}
type TodosRes struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Title     string    `json:"title" db:"title"`
	Complete  bool      `json:"complete" db:"complete"`
}
