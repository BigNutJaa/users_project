package wrapper

import (
	"go.uber.org/dig"

	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/delta"
)

type Wrapper struct {
	dig.In  `name:"wrapperDelta"`
	Service service.Service
}
