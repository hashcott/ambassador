package repository

import (
	"github.com/fungerouscode/go-ambassador/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(firstName string, lastName string, email string, password string) (models.User, error) {
	newUser := models.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		Password:     password,
		IsAmbassador: false,
	}

	r.db.Create(&newUser)

	return newUser, nil
}
