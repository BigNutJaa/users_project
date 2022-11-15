package wrapper

import (
	"go.uber.org/dig"

	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/moving"
)

type Wrapper struct {
	dig.In  `name:"wrapperMoving"`
	Service service.Service
}
