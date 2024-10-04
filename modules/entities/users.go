package entities

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type UsersUsecase interface {
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
	GetAllUsers(c *fiber.Ctx, req *GetAllUserReq) ([]UsersAllRes, error)
}

type UsersRepository interface {
	FindOneUser(username string) (*UsersPassport, error)
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
	GetAllUsers(req *GetAllUserReq) ([]UsersAllRes, error)
}

type UsersRegisterReq struct {
	ID       int64     `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Email    string    `json:"email" db:"email"`
	Role     string    `json:"role" db:"role"`
	CreateAt time.Time `json:"create_at" db:"create_at"`
}

type UsersRegisterRes struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

type UsersAllRes struct {
	ID       int64     `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Email    string    `json:"email" db:"email"`
	CreateAt time.Time `json:"create_at" db:"create_at"`
}

type UserRole string

const (
	RoleAdmin   UserRole = "admin"
	RoleUser    UserRole = "user"
	RoleManager UserRole = "manager"
)

type GetAllUserReq struct {
	Username string `json:"username" db:"username"`
}
