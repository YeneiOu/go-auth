package servers

import (
	"clean-arc/helpers"
	_authHttp "clean-arc/modules/auth/controllers"
	_authRepository "clean-arc/modules/auth/repositories"
	_authUsecase "clean-arc/modules/auth/usecases"

	_usersHttp "clean-arc/modules/users/controllers"
	_usersRepository "clean-arc/modules/users/repositories"
	_usersUsecase "clean-arc/modules/users/usecases"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {

	// Group a version
	v1 := s.App.Group("/v1")

	//* Users group
	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepository(s.Db)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepository)
	_usersHttp.NewUsersController(usersGroup, usersUsecase)

	authGroup := v1.Group("/auth")
	authRepository := _authRepository.NewAuthRepository(s.Db)
	authUsecase := _authUsecase.NewAuthUsecase(authRepository, usersRepository)
	_authHttp.NewAuthController(authGroup, s.Cfg, authUsecase)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return helpers.RespondWithJSON(c, fiber.ErrInternalServerError.Message, fiber.ErrInternalServerError.Code, "error, end point not found", nil)
	})

	return nil
}
