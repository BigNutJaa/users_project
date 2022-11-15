package wrapper

import (
	"context"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/stock"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Get(ctx context.Context, input *model.FitterReadStock) (*model.ReadResponseStock, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Stock.Get")
	defer sp.Finish()

	id, err := wrp.Service.Get(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
