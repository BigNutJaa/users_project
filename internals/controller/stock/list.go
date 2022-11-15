package stock

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/stock"
	"github.com/BigNutJaa/user-service/internals/utils"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (c *Controller) List(ctx context.Context, request *apiV1.StockListRequest) (*apiV1.StockListResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.stock.List",
	)
	defer span.Finish()
	pg := &utils.PaginationOptions{
		Page:        request.GetPage(),
		ItemPerPage: request.GetPerPage(),
	}
	span.LogKV("Input Handler", request)
	stockData, err := c.service.List(ctx, &model.FitterListStock{
		Name:   request.GetName(),
		Detail: request.GetDetail(),
	}, pg)

	if err != nil {
		return nil, err
	}
	return c.mapListStock(stockData)
}
func (c *Controller) mapListStock(input *utils.Pagination) (*apiV1.StockListResponse, error) {
	items, ok := input.Items.([]*model.ListResponseStock)
	if !ok {
		return nil, errors.New("Error Struct")
	}
	respone := &apiV1.StockListResponse{
		Size:       input.Size,
		Total:      input.Total,
		TotalPages: input.TotalPages,
	}
	for _, item := range items {
		respone.Item = append(respone.Item, &apiV1.StockListItem{
			Name:   item.Name,
			Detail: item.Detail,
			Qty:    item.Qty,
			Id:     item.Id,
		})
	}
	return respone, nil
}
