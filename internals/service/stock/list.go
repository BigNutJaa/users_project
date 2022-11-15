package stock

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/stock"
)

func (s *StockService) Get(ctx context.Context, request *model.FitterReadStock) (*model.ReadResponseStock, error) {
	makeFilter := s.makeFilterStock(request)
	stock := &entity.Stock{}

	err := s.repository.Find(makeFilter, stock)

	return &model.ReadResponseStock{
		Name:   stock.Name,
		Detail: stock.Detail,
		Qty:    stock.Qty,
		Id:     int32(stock.ID),
	}, err
}

func (s *StockService) makeFilterStock(filters *model.FitterReadStock) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.Name) > 0 {
		output["name"] = filters.Name
	}
	if len(filters.Detail) > 0 {
		output["detail"] = filters.Detail
	}
	if filters.Id > 0 {
		output["id"] = filters.Id
	}
	return output

}
