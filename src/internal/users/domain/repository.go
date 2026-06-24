package domain

import (
	"LogiredAPIWeb/src/internal/users/domain/entities"
	"database/sql"
)

type UserRepository interface {
	CreateUser(entities.User) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetUserByID(IdUsuario int32) (entities.User, error)
	UpdateUser(IdUsuario int32, user entities.User) error
	RessetPassword(email string, password string) error
	UpdatePassword(IdUsuario int32, password string) error
	DeleteUser(IdUsuario int32) error
	CreateUserTx(tx *sql.Tx, u entities.User) (entities.User, error)
	BeginTx() (*sql.Tx, error)
}
