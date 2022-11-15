package alpha

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/alpha"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.AlphaCreateRequest) (*apiV1.AlphaCreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.alpha.Create",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	id, err := c.service.Create(ctx, &model.Request{
		MovieName: request.GetMovieName(),
		Date:      request.GetDate(),
		Time:      request.GetTime(),
		CinemaNo:  request.GetCinemaNo(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.AlphaCreateResponse{Id: int32(id)}, nil
}
