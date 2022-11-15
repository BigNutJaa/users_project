package stock

import (
	"context"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/stock"
)

func (s *StockService) Create(ctx context.Context, request *model.Request) (int, error) {

	input := &entity.Stock{
		Name:   request.Name,
		Detail: request.Detail,
		Qty:    request.Qty,
	}

	err := s.repository.Create(input)

	//sp.LogKV("Repository result  :", err)

	return input.ID, err
}
