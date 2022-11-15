package alpha

import (
	"github.com/BigNutJaa/user-service/internals/repository/postgres"
)

type AlphaService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &AlphaService{
		repository: r,
	}
}
