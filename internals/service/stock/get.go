package stock

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/stock"
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
