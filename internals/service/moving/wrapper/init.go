package wrapper

import (
	"go.uber.org/dig"

	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/movie"
)

type Wrapper struct {
	dig.In  `name:"wrapperMovie"`
	Service service.Service
}
