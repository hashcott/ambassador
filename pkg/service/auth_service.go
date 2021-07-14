package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fungerouscode/go-ambassador/models"
	"github.com/fungerouscode/go-ambassador/pkg/repository"
)

const (
	salt       = "jsdsgdahfdasvbd77et63y"
	tokenTTL   = 12 * time.Hour
	signingKey = "dbshgdsvbdfs7363"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

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

func (s *AuthService) GenerateToken(email string, password string) (string, error) {
	// Get user from db
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}
