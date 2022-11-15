package container

import (
	controllerAlpha "github.com/BigNutJaa/user-service/internals/controller/alpha"
	controllerDelta "github.com/BigNutJaa/user-service/internals/controller/delta"
	controllerMovie "github.com/BigNutJaa/user-service/internals/controller/movie"
	controllerMoving "github.com/BigNutJaa/user-service/internals/controller/moving"
	"github.com/BigNutJaa/user-service/internals/service/alpha"
	"github.com/BigNutJaa/user-service/internals/service/delta"
	"github.com/BigNutJaa/user-service/internals/service/movie"
	"github.com/BigNutJaa/user-service/internals/service/moving"
	"github.com/BigNutJaa/user-service/internals/service/stock"
	"net/http"

	controllerUsers "github.com/BigNutJaa/user-service/internals/controller/users"
	"github.com/BigNutJaa/user-service/internals/service/users"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/robowealth-mutual-fund/shared-utility/validator"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"github.com/BigNutJaa/user-service/internals/config"
	"github.com/BigNutJaa/user-service/internals/controller"
	controllerCategory "github.com/BigNutJaa/user-service/internals/controller/category"
	controllerProduct "github.com/BigNutJaa/user-service/internals/controller/product"
	controllerStock "github.com/BigNutJaa/user-service/internals/controller/stock"
	warehouseController "github.com/BigNutJaa/user-service/internals/controller/warehouse"
	"github.com/BigNutJaa/user-service/internals/infrastructure/database"
	grpcServer "github.com/BigNutJaa/user-service/internals/infrastructure/grpcServer"
	httpServer "github.com/BigNutJaa/user-service/internals/infrastructure/httpServer"
	"github.com/BigNutJaa/user-service/internals/infrastructure/jaeger"
	"github.com/BigNutJaa/user-service/internals/repository/postgres"
	"github.com/BigNutJaa/user-service/internals/service/category"
	serviceProduct "github.com/BigNutJaa/user-service/internals/service/product"
	warehouseService "github.com/BigNutJaa/user-service/internals/service/warehouse"
	"github.com/BigNutJaa/user-service/internals/utils"
	"github.com/BigNutJaa/user-service/internals/utils/logrus"
)

type Container struct {
	container *dig.Container
}

func (c *Container) Configure() error {

	servicesConstructors := []interface{}{
		config.NewConfiguration,
		grpcServer.NewServer,
		database.NewServerBase,
		http.NewServeMux,
		httpServer.NewServer,
		runtime.NewServeMux,
		jaeger.NewJaeger,
		logrus.NewLog,
		controller.NewHealthZController,
		controller.NewPingPongController,
		validator.NewCustomValidator,
		controllerProduct.NewController,
		serviceProduct.NewService,
		postgres.NewRepository,
		utils.NewUtils,
		utils.NewCustomValidator,
		category.NewService,
		warehouseService.NewService,
		users.NewService,
		controllerUsers.NewController,
		stock.NewService,
		controllerStock.NewController,
		movie.NewService,
		moving.NewService,
		delta.NewService,
		alpha.NewService,
		controllerAlpha.NewController,
		controllerMovie.NewController,
		controllerMoving.NewController,
		controllerDelta.NewController,
		controllerCategory.NewController,
		warehouseController.NewController,
	}
	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}
	appConfig := config.NewConfiguration()
	jaeger.NewJaeger(appConfig)
	logrus.NewLog()
	return nil
}

func (c *Container) Start() error {
	log.Info("Start Container")
	if err := c.container.Invoke(func(s *grpcServer.Server, h *httpServer.Server, v *validator.CustomValidator) {
		go func() {
			_ = h.Start()
		}()
		s.Start()

	}); err != nil {
		log.Errorf("%s", err)

		return err
	}

	return nil
}

// MigrateDB ...
func (c *Container) MigrateDB() error {
	log.Info("Start Container DB")

	if err := c.container.Invoke(func(d *database.DB) {
		d.MigrateDB()
	}); err != nil {
		return err
	}

	return nil
}

func NewContainer() (*Container, error) {
	d := dig.New()

	container := &Container{
		container: d,
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	return container, nil
}
