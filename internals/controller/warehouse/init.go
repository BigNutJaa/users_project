package warehouse

import (
	"github.com/BigNutJaa/user-service/internals/service/warehouse/wrapper"
)

type Controller struct {
	service wrapper.Wrapper
}

func NewController(s wrapper.Wrapper) *Controller {
	return &Controller{
		service: s,
	}
}
