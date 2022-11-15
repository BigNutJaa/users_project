package warehouse

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/warehouse"
)

func (s *WarehouseService) Create(ctx context.Context, request *model.Request) (int, error) {

	input := &entity.Warehouse{
		Name:   request.Name,
		Detail: request.Detail,
	}

	err := s.repository.Create(input)

	//sp.LogKV("Repository result  :", err)

	return input.ID, err
}
