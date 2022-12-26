package wrapper

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/users"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Delete(ctx context.Context, input *model.FitterDeleteUsers) (*model.DeleteResponseUsers, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.User.Delete")
	defer sp.Finish()

	id, err := wrp.Service.Delete(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
