package movie

import (
	"github.com/BigNutJaa/user-service/internals/repository/postgres"
)

type MovieService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &MovieService{
		repository: r,
	}
}
