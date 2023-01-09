package token

import (
	"context"

	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/token"
)

func (s *LoginService) Delete(ctx context.Context, request *model.FitterDeleteToken) (*model.DeleteResponseToken, error) {
	makeFilter := s.makeFilterTokenDelete(request)
	tokenDelete := &entity.Token{}

	err := s.repository.Delete(makeFilter, tokenDelete)

	return &model.DeleteResponseToken{
		Id: int32(tokenDelete.ID),
	}, err
}

func (s *LoginService) makeFilterTokenDelete(filters *model.FitterDeleteToken) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if filters.Id > 0 {
		output["id"] = filters.Id
	}

	return output

}
