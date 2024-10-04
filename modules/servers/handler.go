package servers

import (
	_authHttp "clean-arc/modules/auth/controllers"
	_authRepository "clean-arc/modules/auth/repositories"
	_authUsecase "clean-arc/modules/auth/usecases"
	_todosHttp "clean-arc/modules/todos/controllers"
	_todosRepository "clean-arc/modules/todos/repositories"
	_todosUsecase "clean-arc/modules/todos/usecases"
	_usersHttp "clean-arc/modules/users/controllers"
	_usersRepository "clean-arc/modules/users/repositories"
	_usersUsecase "clean-arc/modules/users/usecases"
	"clean-arc/pkg/utils"
)

func (s *Server) MapHandlers() error {

	// Apply CORS middleware to the application
	s.App.Use(utils.CORS)

	// Group API version 1
	v1 := s.App.Group("/api/v1")

	//* Users group
	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepository(s.Db)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepository)
	_usersHttp.NewUsersController(usersGroup, usersUsecase)

	// Auth group
	authGroup := v1.Group("/auth")
	authRepository := _authRepository.NewAuthRepository(s.Db)
	authUsecase := _authUsecase.NewAuthUsecase(authRepository, usersRepository)
	_authHttp.NewAuthController(authGroup, s.Cfg, authUsecase)

	//todos
	todosGroup := v1.Group("/todos")
	todosRepository := _todosRepository.NewTodosRepository(s.Db, s.RedisClient)
	todosUsecase := _todosUsecase.NewTodosUsecase(todosRepository)
	_todosHttp.NewTodosController(todosGroup, todosUsecase)

	return nil
}
