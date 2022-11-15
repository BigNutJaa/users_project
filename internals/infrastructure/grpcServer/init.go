package grpcserver

import (
	"github.com/BigNutJaa/user-service/internals/config"
	"github.com/BigNutJaa/user-service/internals/controller"
	controllerAlpha "github.com/BigNutJaa/user-service/internals/controller/alpha"
	controllerCategory "github.com/BigNutJaa/user-service/internals/controller/category"
	controllerDelta "github.com/BigNutJaa/user-service/internals/controller/delta"
	controllerMovie "github.com/BigNutJaa/user-service/internals/controller/movie"
	controllerMoving "github.com/BigNutJaa/user-service/internals/controller/moving"
	controllerProduct "github.com/BigNutJaa/user-service/internals/controller/product"
	controllerStock "github.com/BigNutJaa/user-service/internals/controller/stock"
	controllerUsers "github.com/BigNutJaa/user-service/internals/controller/users"
	controllerWarehouse "github.com/BigNutJaa/user-service/internals/controller/warehouse"
	apiV1 "github.com/BigNutJaa/user-service/pkg/api/v1"
	grpcErrors "github.com/robowealth-mutual-fund/shared-utility/grpc_errors"
	validatorUtils "github.com/robowealth-mutual-fund/shared-utility/validator"
	"google.golang.org/grpc"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	Config        config.Configuration
	Server        *grpc.Server
	HealthCtrl    *controller.HealthZController
	PingPongCtrl  *controller.PingPongController
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

// Configure ...
func (s *Server) Configure() {
	grpcHealthV1.RegisterHealthServer(s.Server, s.HealthCtrl)
	apiV1.RegisterPingPongServiceServer(s.Server, s.PingPongCtrl)
	apiV1.RegisterProductServiceServer(s.Server, s.ProductCtrl)
	apiV1.RegisterCategoryServiceServer(s.Server, s.CategoryCtrl)
	apiV1.RegisterWarehouseServiceServer(s.Server, s.WarehouseCtrl)
	apiV1.RegisterUsersServiceServer(s.Server, s.UsersCtrl)
	apiV1.RegisterStockServiceServer(s.Server, s.StockCtrl)
	apiV1.RegisterMovieServiceServer(s.Server, s.MovieCtrl)
	apiV1.RegisterMovingServiceServer(s.Server, s.MovingCtrl)
	apiV1.RegisterDeltaServiceServer(s.Server, s.DeltaCtrl)
	apiV1.RegisterAlphaServiceServer(s.Server, s.AlphaCtrl)
}

func NewServer(
	config config.Configuration,
	healthCtrl *controller.HealthZController,
	pingPongCtrl *controller.PingPongController,
	productCtrl *controllerProduct.Controller,
	categoryCtrl *controllerCategory.Controller,
	warehouseCtrl *controllerWarehouse.Controller,
	usersCtrl *controllerUsers.Controller,
	stockCtrl *controllerStock.Controller,
	movieCtrl *controllerMovie.Controller,
	movingCtrl *controllerMoving.Controller,
	deltaCtrl *controllerDelta.Controller,
	alphaCtrl *controllerAlpha.Controller,
	validator *validatorUtils.CustomValidator,
) *Server {
	option := grpc.ChainUnaryInterceptor(
		grpcErrors.UnaryServerInterceptor(),
		validatorUtils.UnaryServerInterceptor(validator),
	)

	s := &Server{
		Server:        grpc.NewServer(option, grpc.MaxRecvMsgSize(10*10e6), grpc.MaxSendMsgSize(10*10e6)),
		Config:        config,
		HealthCtrl:    healthCtrl,
		PingPongCtrl:  pingPongCtrl,
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

	s.Configure()

	return s
}
