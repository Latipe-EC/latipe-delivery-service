//go:build wireinject
// +build wireinject

package server

import (
	"delivery-service/config"
	"delivery-service/internal/adapter"
	"delivery-service/internal/api"
	"delivery-service/internal/domain/repos"
	"delivery-service/internal/middleware"
	"delivery-service/internal/router"
	"delivery-service/internal/service"
	"delivery-service/pkgs/mongodb"
	restyclient "delivery-service/pkgs/resty"
	"encoding/json"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/wire"
	"time"
)

type Server struct {
	app       *fiber.App
	globalCfg *config.Config
}

func New() (*Server, error) {
	panic(wire.Build(wire.NewSet(
		NewServer,
		config.Set,
		restyclient.Set,
		mongodb.Set,
		repos.Set,
		adapter.Set,
		service.Set,
		api.Set,
		middleware.Set,
		router.Set,
	)))
}

func NewServer(
	cfg *config.Config,
	router *router.RouterHandler) *Server {

	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	prometheus := fiberprometheus.New("latipe-delivery-service")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Initialize default config
	app.Use(logger.New())

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

	api := app.Group("/api")
	v1 := api.Group("/v1")

	router.InitRouter(&v1)

	return &Server{
		globalCfg: cfg,
		app:       app,
	}
}

func (serv Server) App() *fiber.App {
	return serv.app
}

func (serv Server) Config() *config.Config {
	return serv.globalCfg
}
