package repository

import (
	"github.com/fungerouscode/go-ambassador/pkg/repository"
)

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
