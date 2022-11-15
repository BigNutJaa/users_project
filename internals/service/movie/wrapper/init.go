package wrapper

import (
	"go.uber.org/dig"

	service "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/stock"
)

type Wrapper struct {
	dig.In  `name:"wrapperStock"`
	Service service.Service
}
