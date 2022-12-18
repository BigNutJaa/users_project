package users

import (
	"context"

	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/users"
)

func (s *UsersService) Get(ctx context.Context, request *model.FitterReadUsers) (*model.ReadResponseUsers, error) {
	makeFilter := s.makeFilterUsers(request)
	users := &entity.Users{}
	err := s.repository.Find(makeFilter, users)
	decryptPass := StartDecrypt(users.Password)

	return &model.ReadResponseUsers{
		Id:         int32(users.ID),
		User_name:  users.UserName,
		Password:   decryptPass,
		First_name: users.FirstName,
		Last_name:  users.LastName,
		Email:      users.Email,
		Role_code:  users.RoleCode,
	}, err
}

func (s *UsersService) makeFilterUsers(filters *model.FitterReadUsers) (output map[string]interface{}) {
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
	if filters.Id > 0 {
		output["id"] = filters.Id
	}
	return output

}
