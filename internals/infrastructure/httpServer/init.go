package httpServer

import (
	"context"
	controllerAlpha "github.com/BigNutJaa/user-service/internals/controller/alpha"
	controllerDelta "github.com/BigNutJaa/user-service/internals/controller/delta"
	controllerMovie "github.com/BigNutJaa/user-service/internals/controller/movie"
	controllerMoving "github.com/BigNutJaa/user-service/internals/controller/moving"
	controllerStock "github.com/BigNutJaa/user-service/internals/controller/stock"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/BigNutJaa/user-service/internals/config"
	controllerCategory "github.com/BigNutJaa/user-service/internals/controller/category"
	controllerProduct "github.com/BigNutJaa/user-service/internals/controller/product"
	controllerUsers "github.com/BigNutJaa/user-service/internals/controller/users"
	controllerWarehouse "github.com/BigNutJaa/user-service/internals/controller/warehouse"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
	"google.golang.org/grpc"
)

type Server struct {
	Config        config.Configuration
	Server        *runtime.ServeMux
	HttpMux       *http.ServeMux
	ProductCtrl   *controllerProduct.Controller
	CategoryCtrl  *controllerCategory.Controller
	WarehouseCtrl *controllerWarehouse.Controller
	UsersCtrl     *controllerUsers.Controller
	StockCtrl     *controllerStock.Controller
	MovieCtrl     *controllerMovie.Controller
	MovingCtrl    *controllerMoving.Controller
	DeltaCtrl     *controllerDelta.Controller
	AlphaCtrl     *controllerAlpha.Controller
}

func (s *Server) Configure(ctx context.Context, opts []grpc.DialOption) {
	apiV1.RegisterProductServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterCategoryServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterWarehouseServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterUsersServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterStockServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterMovieServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterMovingServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterDeltaServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterAlphaServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)

}

func NewServer(config config.Configuration, rmux *runtime.ServeMux, httpMux *http.ServeMux,
	productCtrl *controllerProduct.Controller,
	categoryCtrl *controllerCategory.Controller,
	warehouseCtrl *controllerWarehouse.Controller,
	usersCtrl *controllerUsers.Controller,
	stockCtrl *controllerStock.Controller,
	movieCtrl *controllerMovie.Controller,
	movingCtrl *controllerMoving.Controller,
	deltaCtrl *controllerDelta.Controller,
	alphaCtrl *controllerAlpha.Controller,
) *Server {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	s := &Server{
		Config:        config,
		Server:        rmux,
		HttpMux:       httpMux,
		ProductCtrl:   productCtrl,
		CategoryCtrl:  categoryCtrl,
		WarehouseCtrl: warehouseCtrl,
		UsersCtrl:     usersCtrl,
		StockCtrl:     stockCtrl,
		MovieCtrl:     movieCtrl,
		MovingCtrl:    movingCtrl,
		DeltaCtrl:     deltaCtrl,
		AlphaCtrl:     alphaCtrl,
	}
	s.Configure(context.Background(), opts)
	return s
}
