package stock

import (
	"github.com/BigNutJaa/user-service/internals/repository/postgres"
)

type StockService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &StockService{
		repository: r,
	}
}
