package moving

import (
	"context"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/moving"
)

func (s *MovingService) Create(ctx context.Context, request *model.Request) (int, error) {
	s.Create2(request)

	input := &entity.Moving{
		MovieName: request.MovieName,
		Date:      request.Date,
		Time:      request.Time,
		CinemaNo:  request.CinemaNo,
	}

	err := s.repository.Create(input)

	return input.ID, err
}

func (s *MovingService) Create2(request *model.Request) (int, error) {

	input2 := &entity.Movie{
		MovieName: request.MovieName,
		Date:      request.Date,
		Time:      request.Time,
		CinemaNo:  request.CinemaNo,
	}

	err := s.repository.Create(input2)

	//sp.LogKV("Repository result  :", err)

	return input2.ID, err
}
