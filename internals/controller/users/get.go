package users

import (
	"context"

	model "github.com/BigNutJaa/users/internals/model/users"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
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
		User_name:  request.GetUserName(),
		First_name: request.GetFirstName(),
		Email:      request.GetEmail(),
		Id:         request.GetId(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.UsersGetResponse{
		UserName:  usersData.User_name,
		Password:  usersData.Password,
		FirstName: usersData.First_name,
		LastName:  usersData.Last_name,
		Email:     usersData.Email,
		RoleCode:  usersData.Role_code,
		Id:        usersData.Id,
	}, nil
}
