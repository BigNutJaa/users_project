package moving

import (
	"github.com/BigNutJaa/user-service/internals/repository/postgres"
	"github.com/BigNutJaa/user-service/internals/service/movie"
)

type MovingService struct {
	repository postgres.Repository
	///
	MovieService movie.Service
}

func NewService(r postgres.Repository, movieService movie.Service) (service Service) {
	return &MovingService{
		repository: r,
		///
		MovieService: movieService,
	}
}
