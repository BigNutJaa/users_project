package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/user-service/internals/service/alpha"
)

type Wrapper struct {
	dig.In  `name:"wrapperAlpha"`
	Service service.Service
}
