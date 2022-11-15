package movie

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/movie"
	"github.com/BigNutJaa/user-service/internals/utils"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (c *Controller) List(ctx context.Context, request *apiV1.MovieListRequest) (*apiV1.MovieListResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.movie.List",
	)
	defer span.Finish()
	pg := &utils.PaginationOptions{
		Page:        request.GetPage(),
		ItemPerPage: request.GetPerPage(),
	}
	span.LogKV("Input Handler", request)
	movieData, err := c.service.List(ctx, &model.FitterListMovie{
		MovieName: request.GetMovieName(),
		Date:      request.GetDate(),
		Time:      request.GetTime(),
		CinemaNo:  request.GetCinemaNo(),
	}, pg)

	if err != nil {
		return nil, err
	}
	return c.mapListMovie(movieData)
}
func (c *Controller) mapListMovie(input *utils.Pagination) (*apiV1.MovieListResponse, error) {
	items, ok := input.Items.([]*model.ListResponseMovie)
	if !ok {
		return nil, errors.New("Error Struct")
	}
	respone := &apiV1.MovieListResponse{
		Size:       input.Size,
		Total:      input.Total,
		TotalPages: input.TotalPages,
	}
	for _, item := range items {
		respone.Item = append(respone.Item, &apiV1.MovieListItem{
			MovieName: item.MovieName,
			Date:      item.Date,
			Time:      item.Date,
			CinemaNo:  item.CinemaNo,
		})
	}
	return respone, nil
}
