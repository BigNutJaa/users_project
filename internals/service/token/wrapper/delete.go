package wrapper

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/token"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Delete(ctx context.Context, input *model.FitterDeleteToken) (*model.DeleteResponseToken, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Token.Delete")
	defer sp.Finish()

	id, err := wrp.Service.Delete(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
