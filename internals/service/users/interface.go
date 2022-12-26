package users

import (
	"context"
	"github.com/BigNutJaa/users/internals/utils"

	model "github.com/BigNutJaa/users/internals/model/users"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID string, err error)
	Get(ctx context.Context, request *model.FitterReadUsers) (*model.ReadResponseUsers, error)
	List(ctx context.Context, request *model.FitterListUsers, pagination *utils.PaginationOptions) (*utils.Pagination, error)
	Delete(ctx context.Context, input *model.FitterDeleteUsers) (*model.DeleteResponseUsers, error)
	Update(ctx context.Context, request *model.FitterUpdateUsers) (*model.UpdateResponseUsers, error)
}
