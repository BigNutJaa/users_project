package delta

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
)

type DeltaService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &DeltaService{
		repository: r,
	}
}
