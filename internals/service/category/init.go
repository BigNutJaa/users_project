package category

import (
	"github.com/BigNutJaa/user-service/internals/repository/postgres"
)

type CategoryService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &CategoryService{
		repository: r,
	}
}
