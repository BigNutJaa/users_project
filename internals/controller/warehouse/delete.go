package warehouse

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/warehouse"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func (c *Controller) Delete(ctx context.Context, request *apiV1.WarehouseDeleteRequest) (*apiV1.WarehouseDeleteResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.warehouse.Delete",
	)
	defer span.Finish()
	logrus.Info("warehouse")
	logrus.Error("warehouse")
	span.LogKV("Input Handler", request)
	warehouseData, err := c.service.Delete(ctx, &model.FitterDeleteWarehouse{
		Name:   request.GetName(),
		Detail: request.GetDetail(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.WarehouseDeleteResponse{
		Name:   warehouseData.Name,
		Detail: warehouseData.Detail,
		Id:     warehouseData.Id,
	}, nil
}
