package repository

import (
	"github.com/fungerouscode/go-ambassador/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(firstName string, lastName string, email string, password string) (models.User, error)
	GetUser(email string, password string) (models.User, error)
	GetUserById(userID uint) (models.User, error)
	UpdateInfo(id uint, firstName string, lastName string, email string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
