//go:build wireinject
// +build wireinject

package server

import (
	"delivery-service/config"
	"delivery-service/internal/adapter"
	"delivery-service/internal/api"
	"delivery-service/internal/domain/repos"
	"delivery-service/internal/grpc-service/interceptor"
	"delivery-service/internal/grpc-service/protobuf"
	"delivery-service/internal/grpc-service/protobuf/deliveryGrpc"
	"delivery-service/internal/middleware"
	"delivery-service/internal/publisher"
	"delivery-service/internal/router"
	"delivery-service/internal/service"
	"delivery-service/internal/subscribers"
	grpcclient "delivery-service/pkgs/grpc"
	healthService "delivery-service/pkgs/healthservice"
	"delivery-service/pkgs/mongodb"
	"delivery-service/pkgs/rabbitclient"
	restyclient "delivery-service/pkgs/resty"
	"encoding/json"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	recoverFiber "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
	"github.com/hellofresh/health-go/v5"
	"google.golang.org/grpc"
	"time"
)

type Server struct {
	globalCfg   *config.Config
	app         *fiber.App
	grpcServ    *grpc.Server
	purchaseSub *subscribers.PurchaseCreatedSub
}

func New() (*Server, error) {
	panic(wire.Build(wire.NewSet(
		NewServer,
		config.Set,
		restyclient.Set,
		grpcclient.Set,
		rabbitclient.Set,
		mongodb.Set,
		publisher.Set,
		subscribers.Set,
		repos.Set,
		adapter.Set,
		service.Set,
		api.Set,
		protobuf.Set,
		interceptor.Set,
		middleware.Set,
		router.Set,
	)))
}

func NewServer(
	cfg *config.Config,
	router *router.RouterHandler,
	deliServ deliveryGrpc.DeliveryServiceServer,
	purchaseSub *subscribers.PurchaseCreatedSub,
	grpcInterceptor *interceptor.GrpcInterceptor) *Server {

	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	// Initialize default config
	app.Use(logger.New())

	recoverConfig := recoverFiber.ConfigDefault
	app.Use(recoverFiber.New(recoverConfig))

	app.Get("", func(ctx *fiber.Ctx) error {
		s := struct {
			Message string `json:"message"`
			Version string `json:"version"`
		}{
			Message: "delivery service was developed by TienDat",
			Version: "v1.0.1",
		}
		return ctx.JSON(s)
	})

	//providing basic authentication for metrics endpoints
	basicAuthConfig := basicauth.Config{
		Users: map[string]string{
			cfg.Metrics.Username: cfg.Metrics.Password,
		},
	}
	//swagger
	app.Get("/swagger/*", basicauth.New(basicAuthConfig), swagger.HandlerDefault) // default

	// Fiber prometheus
	prometheus := fiberprometheus.New("latipe-delivery-service")
	prometheus.RegisterAt(app, cfg.Metrics.MetricsURL, basicauth.New(basicAuthConfig))
	app.Use(prometheus.Middleware)

	// Healthcheck
	h, _ := healthService.NewHealthCheckService(cfg)
	app.Get("/health", basicauth.New(basicAuthConfig), adaptor.HTTPHandlerFunc(h.HandlerFunc))
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/liveness",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			result := h.Measure(c.Context())
			return result.Status == health.StatusOK
		},
		ReadinessEndpoint: "/readiness",
	}))

	//fiber dashboard
	app.Get(cfg.Metrics.FiberDashboard, basicauth.New(basicAuthConfig),
		monitor.New(monitor.Config{Title: "Delivery Services Metrics Page (Fiber)"}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	router.InitRouter(&v1)
	//grpc
	grpcServ := grpc.NewServer(
		grpc.UnaryInterceptor(grpcInterceptor.MiddlewareUnaryRequest),
	)
	deliveryGrpc.RegisterDeliveryServiceServer(grpcServ, deliServ)
	return &Server{
		globalCfg:   cfg,
		app:         app,
		purchaseSub: purchaseSub,
		grpcServ:    grpcServ,
	}
}

func (serv Server) PurchaseCreatedSub() *subscribers.PurchaseCreatedSub {
	return serv.purchaseSub
}

func (serv Server) App() *fiber.App {
	return serv.app
}

func (serv Server) Config() *config.Config {
	return serv.globalCfg
}

func (serv Server) DeliServ() *grpc.Server {
	return serv.grpcServ
}
