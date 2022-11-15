package wrapper

import (
	"context"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/stock"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Update(ctx context.Context, input *model.FitterUpdateStock) (*model.UpdateResponseStock, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Stock.Update")
	defer sp.Finish()

	id, err := wrp.Service.Update(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
