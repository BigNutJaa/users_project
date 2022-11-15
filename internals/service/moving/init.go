package movie

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
)

type MovieService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &MovieService{
		repository: r,
	}
}
