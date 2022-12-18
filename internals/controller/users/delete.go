package users

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/users"

	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Delete(ctx context.Context, request *apiV1.UsersDeleteRequest) (*apiV1.UsersDeleteResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.stock.Delete",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	stockData, err := c.service.Delete(ctx, &model.FitterDeleteStock{
		Id: request.GetId(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.StockDeleteResponse{
		Name:   stockData.Name,
		Detail: stockData.Detail,
		Qty:    stockData.Qty,
		Id:     stockData.Id,
	}, nil
}
