package wrapper

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/movie"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Update(ctx context.Context, input *model.FitterUpdateMovie) (*model.UpdateResponseMovie, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Movie.Update")
	defer sp.Finish()

	id, err := wrp.Service.Update(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
