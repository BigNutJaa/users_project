package category

import (
	"context"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/category"
)

func (s *CategoryService) Patch(ctx context.Context, request *model.FitterUpdateCategory) (*model.UpdateResponseCategory, error) {
	makeFilter := s.makeFilterCategory(request)
	categoryUpdate := &entity.Category{
		Name:   request.NameUpdate,
		Detail: request.DetailUpdate,
		ID:     int(request.Id),
	}

	err := s.repository.Update(makeFilter, categoryUpdate)

	return &model.UpdateResponseCategory{
		Name:   categoryUpdate.Name,
		Detail: categoryUpdate.Detail,
		Id:     int32(categoryUpdate.ID),
	}, err
}

func (s *CategoryService) makeFilterCategory(filters *model.FitterUpdateCategory) (output map[string]interface{}) {
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
