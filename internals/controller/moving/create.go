package movie

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/movie"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.MovieCreateRequest) (*apiV1.MovieCreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.movie.Create",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	id, err := c.service.Create(ctx, &model.Request{
		MovieName: request.GetMovieName(),
		Date:      request.GetDate(),
		Time:      request.GetTime(),
		CinemaNo:  request.GetCinemaNo(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.MovieCreateResponse{Id: int32(id)}, nil
}
