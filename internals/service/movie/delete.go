package movie

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/model/movie"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/movie"
)

func (s *MovieService) Delete(ctx context.Context, request *movie.FitterDeleteMovie) (*model.DeleteResponseMovie, error) {
	makeFilter := s.makeFilterMovieDelete(request)
	movieDelete := &entity.Movie{}

	err := s.repository.Delete(makeFilter, movieDelete)

	return &model.DeleteResponseMovie{
		Id: int32(movieDelete.ID),
	}, err
}

func (s *MovieService) makeFilterMovieDelete(filters *model.FitterDeleteMovie) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if filters.Id > 0 {
		output["id"] = filters.Id
	}
	return output

}
