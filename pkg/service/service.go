package service

import (
	registration "TBCC_RegistrationPJ"
	"TBCC_RegistrationPJ/pkg/repository"
)

type Authorization interface {
	CreateUser(user registration.User) (int, error)
	GenerateToken(Email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
