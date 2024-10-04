package usecases

import (
	"clean-arc/modules/entities"
	"clean-arc/pkg/utils"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type usersUse struct {
	UsersRepo entities.UsersRepository
}

// Constructor
func NewUsersUsecase(usersRepo entities.UsersRepository) entities.UsersUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
	}
}

func (u *usersUse) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	sameUser, err := u.UsersRepo.FindOneUser(req.Username)
	if err != nil {
		return nil, err
	}

	if sameUser != nil {
		return nil, errors.New("user with the same username already exists")
	}
	// Hash a password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	input := entities.UsersRegisterReq{
		Username: req.Username,
		Password: string(hashed),
		Email:    req.Email,
		Role:     "user",
		CreateAt: time.Now(),
	}

	fmt.Println("input ", input)
	user, err := u.UsersRepo.Register(&input)

	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *usersUse) GetAllUsers(c *fiber.Ctx, req *entities.GetAllUserReq) ([]entities.UsersAllRes, error) {

	user, err := utils.BindingUsername(c)
	fmt.Printf("user %v", user)
	allUsers, err := u.UsersRepo.GetAllUsers(req)

	if err != nil {
		return nil, err
	}

	return allUsers, nil
}
