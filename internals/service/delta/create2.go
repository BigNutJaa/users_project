package delta

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/delta"
)

func (s *DeltaService) Create2(ctx context.Context, request *model.Request) (int, error) {

	go func() {
		for i := 1; i < 2000; i++ {
			input := &entity.Delta{
				MovieName: request.MovieName,
				Date:      request.Date,
				Time:      request.Time,
				CinemaNo:  request.CinemaNo,
			}
			s.repository.Create(input)
		}
	}()

	return 1, nil
}
