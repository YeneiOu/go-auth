package entities

import (
	"clean-arc/configs"
	"github.com/golang-jwt/jwt/v4"
)

type AuthRepository interface {
	SignUsersAccessToken(req *UsersPassport) (string, error)
}

type AuthUsecase interface {
	Login(cfg *configs.Configs, req *UsersCredentials) (*UsersLoginRes, error)
}

type UsersCredentials struct {
	Username string `json:"username" db:"username" form:"username"`
	Password string `json:"password" db:"password" form:"password"`
}

type UsersPassport struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	UserRole string `json:"role" db:"role" default:"user"`
}

type UsersClaims struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	UserRole string `json:"role"`
	jwt.RegisteredClaims
}

type UsersLoginRes struct {
	AccessToken string `json:"access_token"`
}
