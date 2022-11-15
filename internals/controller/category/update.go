package category

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/BigNutJaa/user-service/internals/model/category"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
)

func (c *Controller) Update(ctx context.Context, request *apiV1.CategoryUpdateRequest) (*apiV1.CategoryUpdateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.category.Patch",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	categoryData, err := c.service.Patch(ctx, &model.FitterUpdateCategory{
		Name:         request.GetName(),
		Detail:       request.GetDetail(),
		Id:           request.GetId(),
		NameUpdate:   request.GetNameUpdate(),
		DetailUpdate: request.GetDetailUpdate(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.CategoryUpdateResponse{
		Name:   categoryData.Name,
		Detail: categoryData.Detail,
		Id:     categoryData.Id,
	}, nil
}
