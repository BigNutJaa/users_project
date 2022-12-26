package users

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/users"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Delete(ctx context.Context, request *apiV1.UsersDeleteRequest) (*apiV1.UsersDeleteResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.users.Delete",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	stockData, err := c.service.Delete(ctx, &model.FitterDeleteUsers{
		Id: request.GetId(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.UsersDeleteResponse{
		UserName:  stockData.User_name,
		FirstName: stockData.First_name,
		LastName:  stockData.Last_name,
		Id:        stockData.Id,
	}, nil
}
