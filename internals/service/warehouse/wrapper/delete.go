package wrapper

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/warehouse"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Delete(ctx context.Context, input *model.FitterDeleteWarehouse) (*model.DeleteResponseWarehouse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Warehouse.Delete")
	defer sp.Finish()

	id, err := wrp.Service.Delete(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
