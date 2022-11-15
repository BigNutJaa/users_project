package movie

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/BigNutJaa/user-service/internals/model/movie"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
)

func (c *Controller) Update(ctx context.Context, request *apiV1.MovieUpdateRequest) (*apiV1.MovieUpdateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.movie.Update",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	movieData, err := c.service.Update(ctx, &model.FitterUpdateMovie{
		MovieName:      request.GetMovieName(),
		Date:           request.GetDate(),
		Time:           request.GetTime(),
		CinemaNo:       request.GetCinemaNo(),
		DateUpdate:     request.GetDateUpdate(),
		TimeUpdate:     request.GetTimeUpdate(),
		CinemaNoUpdate: request.GetCinemaNoUpdate(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.MovieUpdateResponse{
		MovieName: movieData.MovieName,
		Date:      movieData.Date,
		Time:      movieData.Time,
		CinemaNo:  movieData.CinemaNo,
		Id:        movieData.Id,
	}, nil
}
