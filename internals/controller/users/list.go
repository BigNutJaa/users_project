package users

import (
	"context"
	"errors"
	"github.com/BigNutJaa/users/internals/utils"

	model "github.com/BigNutJaa/users/internals/model/users"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) List(ctx context.Context, request *apiV1.UsersListRequest) (*apiV1.UsersListResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.Users.List",
	)
	defer span.Finish()
	pg := &utils.PaginationOptions{
		Page:        request.GetPage(),
		ItemPerPage: request.GetPerPage(),
	}
	span.LogKV("Input Handler", request)
	usersData, err := c.service.List(ctx, &model.FitterListUsers{
		User_name:  request.GetUserName(),
		First_name: request.GetFirstName(),
		Email:      request.GetEmail(),
	}, pg)

	if err != nil {
		return nil, err
	}
	return c.mapListUser(usersData)
}
func (c *Controller) mapListUser(input *utils.Pagination) (*apiV1.UsersListResponse, error) {
	items, ok := input.Items.([]*model.ListResponseUsers)
	if !ok {
		return nil, errors.New("Error Struct")
	}
	respone := &apiV1.UsersListResponse{
		Size:       input.Size,
		Total:      input.Total,
		TotalPages: input.TotalPages,
	}
	for _, item := range items {
		respone.Item = append(respone.Item, &apiV1.UserListItem{
			UserName:  item.User_name,
			FirstName: item.First_name,
			Email:     item.Email,
			Id:        item.Id,
		})
	}
	return respone, nil
}
