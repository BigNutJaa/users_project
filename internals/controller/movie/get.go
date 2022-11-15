package movie

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/BigNutJaa/user-service/internals/model/movie"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
)

func (c *Controller) Get(ctx context.Context, request *apiV1.MovieGetRequest) (*apiV1.MovieGetResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.movie.Get",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	movieData, err := c.service.Get(ctx, &model.FitterReadMovie{
		MovieName: request.GetMovieName(),
		Date:      request.GetDate(),
		Time:      request.GetTime(),
		CinemaNo:  request.GetCinemaNo(),
		Id:        request.GetId(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.MovieGetResponse{
		MovieName: movieData.MovieName,
		Date:      movieData.Date,
		Time:      movieData.Time,
		CinemaNo:  movieData.CinemaNo,
	}, nil
}
