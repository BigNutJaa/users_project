package stock

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/stock"
)

func (s *StockService) Update(ctx context.Context, request *model.FitterUpdateStock) (*model.UpdateResponseStock, error) {
	makeFilter := s.makeFilterStockUpdate(request)
	stockUpdate := &entity.Stock{
		Qty: request.QtyUpdate,
		ID:  int(request.Id),
	}

	err := s.repository.Update(makeFilter, stockUpdate)

	return &model.UpdateResponseStock{
		Name:   stockUpdate.Name,
		Detail: stockUpdate.Detail,
		Qty:    stockUpdate.Qty,
		Id:     int32(stockUpdate.ID),
	}, err
}

func (s *StockService) makeFilterStockUpdate(filters *model.FitterUpdateStock) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.Name) > 0 {
		output["name"] = filters.Name
	}
	if len(filters.Detail) > 0 {
		output["detail"] = filters.Detail
	}
	if filters.Qty > 0 {
		output["qty"] = filters.Qty
	}
	if filters.Id > 0 {
		output["id"] = filters.Id
	}
	return output

}
