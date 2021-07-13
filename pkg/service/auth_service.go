package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/fungerouscode/go-ambassador/models"
	"github.com/fungerouscode/go-ambassador/pkg/repository"
)

const salt = "jsdsgdahfdasvbd77et63y"

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(firstName string, lastName string, email string, password string) (models.User, error) {
	password = generatePasswordHash(password)
	return s.repo.CreateUser(firstName, lastName, email, password)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
