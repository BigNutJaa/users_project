package movie

import (
	"context"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/movie"
)

func (s *MovieService) Get(ctx context.Context, request *model.FitterReadMovie) (*model.ReadResponseMovie, error) {
	makeFilter := s.makeFilterMovie(request)
	movie := &entity.Movie{}

	err := s.repository.Find(makeFilter, movie)

	return &model.ReadResponseMovie{
		MovieName: movie.MovieName,
		Date:      movie.Date,
		Time:      movie.Time,
		CinemaNo:  movie.CinemaNo,
		Id:        int32(movie.ID),
	}, err
}

func (s *MovieService) makeFilterMovie(filters *model.FitterReadMovie) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.MovieName) > 0 {
		output["movieName"] = filters.MovieName
	}
	if len(filters.Date) > 0 {
		output["date"] = filters.Date
	}
	if len(filters.Time) > 0 {
		output["time"] = filters.Time
	}
	if len(filters.CinemaNo) > 0 {
		output["cinemaNo"] = filters.CinemaNo
	}
	if filters.Id > 0 {
		output["id"] = filters.Id
	}
	return output

}
