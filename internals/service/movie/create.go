package movie

import (
	"context"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/movie"
)

func (s *MovieService) Create(ctx context.Context, request *model.Request) (int, error) {

	input := &entity.Movie{
		MovieName: request.MovieName,
		Date:      request.Date,
		Time:      request.Time,
		CinemaNo:  request.CinemaNo,
	}

	err := s.repository.Create(input)

	//sp.LogKV("Repository result  :", err)

	return input.ID, err
}
