package wrapper

import (
	"context"
	"github.com/BigNutJaa/users/internals/utils"

	model "github.com/BigNutJaa/users/internals/model/users"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) List(ctx context.Context, input *model.FitterListUsers, pagination *utils.PaginationOptions) (*utils.Pagination, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Users.List")
	defer sp.Finish()

	id, err := wrp.Service.List(ctx, input, pagination)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
