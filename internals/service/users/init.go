package users

import (
	"github.com/BigNutJaa/users/internals/repository/postgres"
)

type UsersService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &UsersService{
		repository: r,
	}
}
