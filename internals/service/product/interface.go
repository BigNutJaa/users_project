package wrapper

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/product"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
}
