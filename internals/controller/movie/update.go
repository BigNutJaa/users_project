package stock

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/model/stock"
	apiV1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"
)

func (c *Controller) Update(ctx context.Context, request *apiV1.StockUpdateRequest) (*apiV1.StockUpdateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.stock.Update",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	stockData, err := c.service.Update(ctx, &model.FitterUpdateStock{
		Name:      request.GetName(),
		Detail:    request.GetDetail(),
		Qty:       request.GetQty(),
		Id:        request.GetId(),
		QtyUpdate: request.GetQtyUpdate(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.StockUpdateResponse{
		Name:   stockData.Name,
		Detail: stockData.Detail,
		Qty:    stockData.Qty,
		Id:     stockData.Id,
	}, nil
}
