package repository

import (
	registration "TBCC_RegistrationPJ"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user registration.User) (int, error)
	CreateCounter() (int, error)
	GetUser(userEmail, password string) (registration.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
