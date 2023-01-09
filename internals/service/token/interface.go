package token

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/token"
)

//go:generate mockery --name=Service
type Service interface {
	//Create(ctx context.Context, input *model.Request) (ID string, err error)
	//Get(ctx context.Context, request *model.FitterReadUsers) (*model.ReadResponseUsers, error)
	//List(ctx context.Context, request *model.FitterListUsers, pagination *utils.PaginationOptions) (*utils.Pagination, error)
	Delete(ctx context.Context, input *model.FitterDeleteToken) (*model.DeleteResponseToken, error)
	//Update(ctx context.Context, request *model.FitterUpdateUsers) (*model.UpdateResponseUsers, error)
}
