package users

import (
	"context"
	"errors"
	"github.com/BigNutJaa/users/internals/utils"

	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/users"
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
			Id:         int32(v.ID),
			User_name:  v.UserName,
			First_name: v.FirstName,
			Email:      v.Email,
		})
	}
	return list
}

func (s *UsersService) makeFilterUsersList(filters *model.FitterListUsers) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.User_name) > 0 {
		output["user_name"] = filters.User_name
	}
	if len(filters.First_name) > 0 {
		output["first_name"] = filters.First_name
	}
	if len(filters.Email) > 0 {
		output["email"] = filters.Email
	}
	return output

}
