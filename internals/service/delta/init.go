package delta

import (
	"github.com/BigNutJaa/user-service/internals/repository/postgres"
)

type DeltaService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &DeltaService{
		repository: r,
	}
}
