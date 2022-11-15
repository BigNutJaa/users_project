package movie

import (
	"context"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/movie"
)

func (s *MovieService) Update(ctx context.Context, request *model.FitterUpdateMovie) (*model.UpdateResponseMovie, error) {
	makeFilter := s.makeFilterMovieUpdate(request)
	movieUpdate := &entity.Movie{
		Time:     request.TimeUpdate,
		Date:     request.DateUpdate,
		CinemaNo: request.CinemaNoUpdate,
		ID:       int(request.Id),
	}

	err := s.repository.Update(makeFilter, movieUpdate)

	return &model.UpdateResponseMovie{
		MovieName: movieUpdate.MovieName,
		Date:      movieUpdate.Date,
		Time:      movieUpdate.Time,
		CinemaNo:  movieUpdate.CinemaNo,
		Id:        int32(movieUpdate.ID),
	}, err
}

func (s *MovieService) makeFilterMovieUpdate(filters *model.FitterUpdateMovie) (output map[string]interface{}) {
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
