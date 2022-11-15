package wrapper

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/category"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Patch(ctx context.Context, input *model.FitterUpdateCategory) (*model.UpdateResponseCategory, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Category.Patch")
	defer sp.Finish()

	id, err := wrp.Service.Patch(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
