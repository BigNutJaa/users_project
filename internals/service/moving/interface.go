package movie

import (
	"context"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/movie"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/utils"

	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/movie"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
	Get(ctx context.Context, request *model.FitterReadMovie) (*model.ReadResponseMovie, error)
	Delete(ctx context.Context, input *movie.FitterDeleteMovie) (*model.DeleteResponseMovie, error)
	Update(ctx context.Context, request *model.FitterUpdateMovie) (*model.UpdateResponseMovie, error)
	List(ctx context.Context, request *model.FitterListMovie, pagination *utils.PaginationOptions) (*utils.Pagination, error)
}
