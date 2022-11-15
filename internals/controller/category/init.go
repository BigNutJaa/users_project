package category

import (
	"github.com/BigNutJaa/user-service/internals/service/category/wrapper"
)

type Controller struct {
	service wrapper.Wrapper
}

func NewController(s wrapper.Wrapper) *Controller {
	return &Controller{
		service: s,
	}
}
