package wrapper

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/model/movie"
	model "github.com/BigNutJaa/user-service/internals/model/movie"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Delete(ctx context.Context, input *movie.FitterDeleteMovie) (*model.DeleteResponseMovie, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Movie.Delete")
	defer sp.Finish()

	id, err := wrp.Service.Delete(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
