package alpha

import (
	"context"
	model "github.com/BigNutJaa/user-service/internals/model/alpha"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
	"github.com/opentracing/opentracing-go"
)

func (c *Controller) Create2(ctx context.Context, request *apiV1.Alpha2CreateRequest) (*apiV1.Alpha2CreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.alpha2.Create2",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	id, err := c.service.Create2(ctx, &model.Request{
		MovieName: request.GetMovieName(),
		Date:      request.GetDate(),
		Time:      request.GetTime(),
		CinemaNo:  request.GetCinemaNo(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.Alpha2CreateResponse{Id: int32(id)}, nil
}
