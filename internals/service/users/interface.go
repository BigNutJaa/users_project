package users

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/utils"

	model "github.com/BigNutJaa/user-service/internals/model/users"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
	Get(ctx context.Context, request *model.FitterReadUsers) (*model.ReadResponseUsers, error)
	List(ctx context.Context, request *model.FitterListUsers, pagination *utils.PaginationOptions) (*utils.Pagination, error)
}
