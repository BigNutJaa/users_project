package wrapper

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/users"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Update(ctx context.Context, input *model.FitterUpdateUsers) (*model.UpdateResponseUsers, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Users.Update")
	defer sp.Finish()

	id, err := wrp.Service.Update(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
