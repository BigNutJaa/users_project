package movie

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/BigNutJaa/user-service/internals/model/movie"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
)

func (c *Controller) Delete(ctx context.Context, request *apiV1.MovieDeleteRequest) (*apiV1.MovieDeleteResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.movie.Delete",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	movieData, err := c.service.Delete(ctx, &model.FitterDeleteMovie{
		Id: request.GetId(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.MovieDeleteResponse{
		Id: movieData.Id,
	}, nil
}
