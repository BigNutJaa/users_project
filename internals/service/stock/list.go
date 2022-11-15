package stock

import (
	"context"
	"github.com/pkg/errors"
	"github.com/BigNutJaa/user-service/internals/utils"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/stock"
)

func (s *StockService) List(ctx context.Context, request *model.FitterListStock, pagination *utils.PaginationOptions) (*utils.Pagination, error) {
	makeFilter := s.makeFilterStockList(request)

	offset, limit := utils.GetOffsetLimit(pagination)

	list, err := s.repository.List("stocks", offset, limit, makeFilter, "created_at DESC", &[]*entity.Stock{})
	if err != nil {
		return nil, err
	}
	items, ok := list.Items.(*[]*entity.Stock)
	if !ok {
		return nil, errors.New("Error")
	}
	list.Items = s.mapListStock(items)

	return list, err
}

func (s *StockService) mapListStock(ent *[]*entity.Stock) []*model.ListResponseStock {

	var list []*model.ListResponseStock
	for _, v := range *ent {
		list = append(list, &model.ListResponseStock{
			Id:     int32(v.ID),
			Name:   v.Name,
			Detail: v.Detail,
			Qty:    v.Qty,
		})
	}
	return list
}

func (s *StockService) makeFilterStockList(filters *model.FitterListStock) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.Name) > 0 {
		output["name"] = filters.Name
	}
	if len(filters.Detail) > 0 {
		output["detail"] = filters.Detail
	}

	return output

}
