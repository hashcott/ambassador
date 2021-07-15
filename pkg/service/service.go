package service

import (
	"github.com/fungerouscode/go-ambassador/models"
	"github.com/fungerouscode/go-ambassador/pkg/repository"
)

type Authorization interface {
	CreateUser(firstName string, lastName string, email string, password string) (models.User, error)
	GenerateToken(email string, password string) (string, error)
	GetUserById(userID uint) (models.User, error)
	ParserToken(token string) (uint, error)
	UpdateInfo(id uint, firstName string, lastName string, email string) (models.User, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
