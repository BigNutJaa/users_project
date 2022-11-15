package alpha

import (
	"github.com/BigNutJaa/user-service/internals/service/alpha/wrapper"
)

type Controller struct {
	service wrapper.Wrapper
}

func NewController(s wrapper.Wrapper) *Controller {
	return &Controller{
		service: s,
	}
}
