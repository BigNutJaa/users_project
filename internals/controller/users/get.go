package users

import (
	"context"

	model "github.com/BigNutJaa/user-service/internals/model/users"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Get(ctx context.Context, request *apiV1.UsersGetRequest) (*apiV1.UsersGetResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.Users.Get",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	usersData, err := c.service.Get(ctx, &model.FitterReadUsers{
		FullName: request.GetFullName(),
		Id:       request.GetId(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.UsersGetResponse{
		FullName:    usersData.FullName,
		Address:     usersData.Address,
		PhoneNumber: usersData.PhoneNumber,
		Gender:      usersData.Gender,
		Id:          usersData.Id,
	}, nil
}
