package wrapper

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/movie"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Get(ctx context.Context, input *model.FitterReadMovie) (*model.ReadResponseMovie, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Movie.Get")
	defer sp.Finish()

	id, err := wrp.Service.Get(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
