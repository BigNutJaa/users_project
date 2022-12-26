package users

import (
	"context"

	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/users"
)

func (s *UsersService) Delete(ctx context.Context, request *model.FitterDeleteUsers) (*model.DeleteResponseUsers, error) {
	makeFilter := s.makeFilterUsersDelete(request)
	usersDelete := &entity.Users{}

	err := s.repository.Delete(makeFilter, usersDelete)

	return &model.DeleteResponseUsers{
		User_name:  usersDelete.UserName,
		First_name: usersDelete.FirstName,
		Last_name:  usersDelete.LastName,
		Id:         int32(usersDelete.ID),
	}, err
}

func (s *UsersService) makeFilterUsersDelete(filters *model.FitterDeleteUsers) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if filters.Id > 0 {
		output["id"] = filters.Id
	}
	if len(filters.User_name) > 0 {
		output["user_name"] = filters.User_name
	}
	return output

}
