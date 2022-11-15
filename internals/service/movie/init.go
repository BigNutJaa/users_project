package stock

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
)

type StockService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &StockService{
		repository: r,
	}
}
