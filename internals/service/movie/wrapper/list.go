package wrapper

import (
	"context"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/stock"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/utils"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) List(ctx context.Context, input *model.FitterListStock, pagination *utils.PaginationOptions) (*utils.Pagination, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Stock.List")
	defer sp.Finish()

	id, err := wrp.Service.List(ctx, input, pagination)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
