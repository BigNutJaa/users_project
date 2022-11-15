package delta

import (
	"github.com/BigNutJaa/user-service/internals/service/delta/wrapper"
)

type Controller struct {
	service wrapper.Wrapper
}

func NewController(s wrapper.Wrapper) *Controller {
	return &Controller{
		service: s,
	}
}
