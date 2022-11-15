package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/user-service/internals/service/product"
)

type Wrapper struct {
	dig.In  `name:"wrapperProduct"`
	Service service.Service
}

func _(service service.Service) service.Service {
	return &Wrapper{
		Service: service,
	}
}
