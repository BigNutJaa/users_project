package moving

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/entity"
	"github.com/BigNutJaa/user-service/internals/model/movie"
	model "github.com/BigNutJaa/user-service/internals/model/moving"
)

func (s *MovingService) Create(ctx context.Context, request *model.Request) (int, error) {
	s.MovieService.Create(ctx, &movie.Request{
		MovieName: request.MovieName,
		Date:      request.Date,
		Time:      request.Time,
		CinemaNo:  request.CinemaNo,
	})

	input := &entity.Moving{
		MovieName: request.MovieName,
		Date:      request.Date,
		Time:      request.Time,
		CinemaNo:  request.CinemaNo,
	}

	err := s.repository.Create(input)

	return input.ID, err
}

//func (s *MovingService) Create2(request *model.Request) (int, error) {
//
//	input2 := &entity.Movie{
//		MovieName: request.MovieName,
//		Date:      request.Date,
//		Time:      request.Time,
//		CinemaNo:  request.CinemaNo,
//	}
//
//	err := s.repository.Create(input2)
//
//	//sp.LogKV("Repository result  :", err)
//
//	return input2.ID, err
//}
