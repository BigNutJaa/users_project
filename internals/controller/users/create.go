package users

import (
	"context"

	model "github.com/BigNutJaa/user-service/internals/model/users"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.UsersCreateRequest) (*apiV1.UsersCreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.category.Create",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	id, err := c.service.Create(ctx, &model.Request{
		FirstName:   request.GetFirstName(),
		LastName:    request.GetLastName(),
		Address:     request.GetAddress(),
		PhoneNumber: request.GetPhoneNumber(),
		Gender:      request.GetGender(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.UsersCreateResponse{Id: int32(id)}, nil
}
