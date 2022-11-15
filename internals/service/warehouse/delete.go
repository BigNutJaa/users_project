package warehouse

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/warehouse"
)

func (s *WarehouseService) Delete(ctx context.Context, request *model.FitterDeleteWarehouse) (*model.DeleteResponseWarehouse, error) {
	makeFilter := s.makeFilterWarehouse(request)
	categoryUpdate := &entity.Warehouse{}

	err := s.repository.Delete(makeFilter, categoryUpdate)

	return &model.DeleteResponseWarehouse{
		Name:   categoryUpdate.Name,
		Detail: categoryUpdate.Detail,
		Id:     int32(categoryUpdate.ID),
	}, err
}

func (s *WarehouseService) makeFilterWarehouse(filters *model.FitterDeleteWarehouse) (output map[string]interface{}) {
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
