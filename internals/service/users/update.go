package users

import (
	"context"
	"golang.org/x/crypto/bcrypt"

	"github.com/BigNutJaa/users/internals/entity"
	model "github.com/BigNutJaa/users/internals/model/users"
)

func (s *UsersService) Update(ctx context.Context, request *model.FitterUpdateUsers) (*model.UpdateResponseUsers, error) {
	makeFilter := s.makeFilterUsersUpdate(request)
	encryptPass, _ := bcrypt.GenerateFromPassword([]byte(request.PasswordUpdate), 10)

	usersUpdate := &entity.Users{
		ID:        int(request.Id),
		UserName:  request.User_name,
		Password:  string(encryptPass),
		FirstName: request.First_nameUpdate,
		LastName:  request.Last_nameUpdate,
		Email:     request.EmailUpdate,
		RoleCode:  request.Role_codeUpdate,
	}

	err := s.repository.Update(makeFilter, usersUpdate)

	return &model.UpdateResponseUsers{
		User_name:  usersUpdate.UserName,
		Password:   usersUpdate.Password,
		First_name: usersUpdate.FirstName,
		Last_name:  usersUpdate.LastName,
		Email:      usersUpdate.Email,
		Role_code:  usersUpdate.RoleCode,
	}, err
}

func (s *UsersService) makeFilterUsersUpdate(filters *model.FitterUpdateUsers) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.User_name) > 0 {
		output["user_name"] = filters.User_name
	}
	if filters.Id > 0 {
		output["id"] = filters.Id
	}
	return output

}
