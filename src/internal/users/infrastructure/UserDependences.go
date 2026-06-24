package infrastructure

import (
	"LogiredAPIWeb/src/core"
	entities_auth "LogiredAPIWeb/src/core/services/auth/domain"
	middleware "LogiredAPIWeb/src/server/middleware"
	"database/sql"

	entities_users "LogiredAPIWeb/src/internal/users/domain"

	"LogiredAPIWeb/src/internal/users/application"
	"LogiredAPIWeb/src/internal/users/infrastructure/controllers"
)

type UserDependencies struct {
	DB             *sql.DB
	Hasher         *core.BcryptHasher
	UserRepo       entities_users.UserRepository
	AuthRepo       entities_auth.AuthRepository
	TokenService   *core.JWTService
}

func NewUserDependencies(
	db 				*sql.DB,
	hasher 			*core.BcryptHasher,
	tokenService	*core.JWTService,
	authRepo 		entities_auth.AuthRepository,
	userRepo 		entities_users.UserRepository,
) *UserDependencies {
	return &UserDependencies{
		DB:             db,
		Hasher:         hasher,
		TokenService:   tokenService,
		AuthRepo:       authRepo,
		UserRepo:       userRepo,
	}
}

func (d *UserDependencies) GetRoutes() *UserRoutes {
	createUserUseCase := application.NewCreateUser(d.UserRepo, d.Hasher)
	getAllUserUseCase := application.NewGetAllUsers(d.UserRepo)
	getUserUseCase := application.NewGetUserByID(d.UserRepo)
	updateUserUseCase := application.NewUpdateUser(d.UserRepo)
	updatePasswordUseCase := application.NewUpdatePassword(d.UserRepo, d.AuthRepo, d.Hasher)
	ressetPasswordUseCase := application.NewRessetPassword(d.UserRepo, d.AuthRepo, d.Hasher)
	deleteUserUseCase := application.NewDeleteUser(d.UserRepo)

	createUserController := controllers.NewCreateUserController(createUserUseCase)
	getUsersController := controllers.NewGetAllUsersController(getAllUserUseCase)
	getUserController := controllers.NewGetByUserIDController(getUserUseCase)
	updateUserController := controllers.NewUpdateUserController(updateUserUseCase)
	updatePasswordController := controllers.NewUpdatePasswordController(updatePasswordUseCase)
	ressetPasswordController := controllers.NewRessetPasswordController(ressetPasswordUseCase)
	deleteUserController := controllers.NewDeleteUserController(deleteUserUseCase)

	authMiddleware := middleware.AuthMiddleware(d.TokenService, d.UserRepo)

	return NewUserRoutes(
		createUserController,
		getUsersController,
		getUserController,
		updateUserController,
		updatePasswordController,
		ressetPasswordController,
		deleteUserController,
		authMiddleware,
	)
}
