package wrapper

import (
	"context"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/stock"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Delete(ctx context.Context, input *model.FitterDeleteStock) (*model.DeleteResponseStock, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Stock.Delete")
	defer sp.Finish()

	id, err := wrp.Service.Delete(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
