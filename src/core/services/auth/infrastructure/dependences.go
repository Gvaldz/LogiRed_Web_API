package infrastructure

import (
	"LogiredAPIWeb/src/core"
	"LogiredAPIWeb/src/core/services/auth/application"
	"LogiredAPIWeb/src/core/services/auth/infrastructure/controllers"
	users_domain "LogiredAPIWeb/src/internal/users/domain"
	"database/sql"
)

type AuthDependencies struct {
	DB       *sql.DB
	Hasher   *core.BcryptHasher
	UserRepo users_domain.UserRepository
}

func NewAuthDependencies(db *sql.DB, hasher *core.BcryptHasher, userRepo users_domain.UserRepository) *AuthDependencies {
	return &AuthDependencies{
		DB:       db,
		Hasher:   hasher,
		UserRepo: userRepo,
	}
}

func (d *AuthDependencies) GetRoutes() *AuthRoutes {
	authRepo := NewAuthRepository(d.DB)
	tokenService := core.NewJWTService()

	loginUC := application.NewLogin(
		authRepo,
		d.UserRepo,
		tokenService,
		d.Hasher,
	)

	loginController := controllers.NewLoginController(loginUC)

	return NewAuthRoutes(loginController)
}
