package users

import (
	"context"
	"errors"
	"github.com/BigNutJaa/user-service/internals/utils"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/users"
)

func (s *UsersService) List(ctx context.Context, request *model.FitterListUsers, pagination *utils.PaginationOptions) (*utils.Pagination, error) {
	makeFilter := s.makeFilterUsersList(request)

	offset, limit := utils.GetOffsetLimit(pagination)

	list, err := s.repository.List("users", offset, limit, makeFilter, "created_at DESC", &[]*entity.Users{})
	if err != nil {
		return nil, err
	}
	items, ok := list.Items.(*[]*entity.Users)
	if !ok {
		return nil, errors.New("Error")
	}
	list.Items = s.mapListUsers(items)

	return list, err
}

func (s *UsersService) mapListUsers(ent *[]*entity.Users) []*model.ListResponseUsers {

	var list []*model.ListResponseUsers
	for _, v := range *ent {
		list = append(list, &model.ListResponseUsers{
			Id:          int32(v.ID),
			FullName:    v.FullName,
			Address:     v.Address,
			PhoneNumber: v.PhoneNumber,
			Gender:      v.Gender,
		})
	}
	return list
}

func (s *UsersService) makeFilterUsersList(filters *model.FitterListUsers) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.FullName) > 0 {
		output["full_name"] = filters.FullName
	}
	if len(filters.Phone_number) > 0 {
		output["phone_number"] = filters.Phone_number
	}
	if len(filters.Gender) > 0 {
		output["gender"] = filters.Gender
	}
	return output

}
