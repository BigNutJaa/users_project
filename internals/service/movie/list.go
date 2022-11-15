package movie

import (
	"context"
	"github.com/BigNutJaa/user-service/internals/utils"
	"github.com/pkg/errors"

	"github.com/BigNutJaa/user-service/internals/entity"
	model "github.com/BigNutJaa/user-service/internals/model/movie"
)

func (s *MovieService) List(ctx context.Context, request *model.FitterListMovie, pagination *utils.PaginationOptions) (*utils.Pagination, error) {
	makeFilter := s.makeFilterMovieList(request)

	offset, limit := utils.GetOffsetLimit(pagination)

	list, err := s.repository.List("movies", offset, limit, makeFilter, "created_at DESC", &[]*entity.Movie{})
	if err != nil {
		return nil, err
	}
	items, ok := list.Items.(*[]*entity.Movie)
	if !ok {
		return nil, errors.New("Error")
	}
	list.Items = s.mapListMovie(items)

	return list, err
}

func (s *MovieService) mapListMovie(ent *[]*entity.Movie) []*model.ListResponseMovie {

	var list []*model.ListResponseMovie
	for _, v := range *ent {
		list = append(list, &model.ListResponseMovie{
			MovieName: v.MovieName,
			Date:      v.Date,
			Time:      v.Time,
			CinemaNo:  v.CinemaNo,
		})
	}
	return list
}

func (s *MovieService) makeFilterMovieList(filters *model.FitterListMovie) (output map[string]interface{}) {
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

	return output

}
