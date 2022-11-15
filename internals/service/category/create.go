package category

import (
	"context"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/category"
)

func (s *CategoryService) Create(ctx context.Context, request *model.Request) (int, error) {

	input := &entity.Category{
		Name:   request.Name,
		Detail: request.Detail,
	}

	err := s.repository.Create(input)

	//sp.LogKV("Repository result  :", err)

	return input.ID, err
}
