package movie

import (
	"github.com/BigNutJaa/user-service/internals/service/movie/wrapper"
)

type Controller struct {
	service wrapper.Wrapper
}

func NewController(s wrapper.Wrapper) *Controller {
	return &Controller{
		service: s,
	}
}
