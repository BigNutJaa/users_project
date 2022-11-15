package stock

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/stock"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
)

func (c *Controller) Get(ctx context.Context, request *apiV1.StockGetRequest) (*apiV1.StockGetResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.stock.Get",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	stockData, err := c.service.Get(ctx, &model.FitterReadStock{
		Name:   request.GetName(),
		Detail: request.GetDetail(),
		Id:     request.GetId(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.StockGetResponse{
		Name:   stockData.Name,
		Detail: stockData.Detail,
		Qty:    stockData.Qty,
		Id:     stockData.Id,
	}, nil
}
