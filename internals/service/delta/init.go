package moving

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
)

type MovingService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &MovingService{
		repository: r,
	}
}
