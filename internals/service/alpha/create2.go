package alpha

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/alpha"
)

func (s *AlphaService) Create2(ctx context.Context, request *model.Request) (int, error) {
	//defer wg.Done()

	input := &entity.Alpha{
		MovieName: request.MovieName,
		Date:      request.Date,
		Time:      request.Time,
		CinemaNo:  request.CinemaNo,
	}

	err := s.repository.Create(input)

	return input.ID, err
}
