package delta

import (
	"context"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/delta"
)

func (s *DeltaService) Create2(ctx context.Context, request *model.Request) (int, error) {

	input := &entity.Delta{
		MovieName: request.MovieName,
		Date:      request.Date,
		Time:      request.Time,
		CinemaNo:  request.CinemaNo,
	}

	err := s.repository.Create(input)

	return input.ID, err
}
