package wrapper

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/movie"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Create(ctx context.Context, input *model.Request) (int, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Movie.Create")
	defer sp.Finish()

	id, err := wrp.Service.Create(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}