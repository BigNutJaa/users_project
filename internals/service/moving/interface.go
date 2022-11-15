package moving

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/moving"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
}
