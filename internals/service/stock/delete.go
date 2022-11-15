package stock

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/model/stock"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/stock"
)

func (s *StockService) Delete(ctx context.Context, request *stock.FitterDeleteStock) (*model.DeleteResponseStock, error) {
	makeFilter := s.makeFilterStockDelete(request)
	stockDelete := &entity.Stock{}

	err := s.repository.Delete(makeFilter, stockDelete)

	return &model.DeleteResponseStock{
		Name:   stockDelete.Name,
		Detail: stockDelete.Detail,
		Qty:    stockDelete.Qty,
		Id:     int32(stockDelete.ID),
	}, err
}

func (s *StockService) makeFilterStockDelete(filters *model.FitterDeleteStock) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if filters.Id > 0 {
		output["id"] = filters.Id
	}
	return output

}
