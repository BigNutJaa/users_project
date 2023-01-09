package token

import (
	"context"
	model "github.com/BigNutJaa/users/internals/model/token"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Delete(ctx context.Context, request *apiV1.TokenDeleteRequest) (*apiV1.TokenDeleteResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.token.Delete",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	tokenData, err := c.service.Delete(ctx, &model.FitterDeleteToken{
		Id: request.GetId(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.TokenDeleteResponse{
		Id: tokenData.Id,
	}, nil
}
