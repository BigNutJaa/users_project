package users

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/users"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Update(ctx context.Context, request *apiV1.UsersUpdateRequest) (*apiV1.UsersUpdateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.users.Update",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	stockData, err := c.service.Update(ctx, &model.FitterUpdateUsers{
		User_name:        request.UserName,
		Id:               request.Id,
		PasswordUpdate:   request.PasswordUpdate,
		First_nameUpdate: request.FirstNameUpdate,
		Last_nameUpdate:  request.LastNameUpdate,
		EmailUpdate:      request.EmailUpdate,
		Role_codeUpdate:  request.RoleCodeUpdate,
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.UsersUpdateResponse{
		UserName:  stockData.User_name,
		Password:  stockData.Password,
		FirstName: stockData.First_name,
		LastName:  stockData.Last_name,
		Email:     stockData.Email,
		RoleCode:  stockData.Role_code,
	}, nil
}
