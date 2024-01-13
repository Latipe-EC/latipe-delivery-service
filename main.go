package main

import (
	"context"
	"delivery-service/adapter/userserv"
	handler "delivery-service/api"
	"delivery-service/config"
	"delivery-service/domain/repos"
	"delivery-service/middleware"
	"delivery-service/pkgs/message"
	"delivery-service/service/deliveryserv"
	"delivery-service/service/shippingserv"
	"encoding/json"
	"fmt"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	fmt.Printf("Init application\n")
	globalCfg, err := config.NewConfig()
	if err != nil {
		panic(err.Error())
	}

	//create connect to mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(globalCfg.Mongodb.ConnectionString).
		SetConnectTimeout(globalCfg.Mongodb.ConnectTimeout*time.Second).
		SetMaxPoolSize(globalCfg.Mongodb.MaxPoolSize))
	db := client.Database("latipe_delivery_db")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	err = message.InitWorkerProducer(globalCfg)
	if err != nil {
		panic(err.Error())
	}

	//create instance fiber
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})
	app.Use(logger.New())
	prometheus := fiberprometheus.New("delivery_service")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Get("/", func(ctx *fiber.Ctx) error {
		s := struct {
			Message string `json:"message"`
			Version string `json:"version"`
		}{
			Message: "Delivery service was developed by TienDat",
			Version: "v0.0.1",
		}
		return ctx.JSON(s)
	})

	//create instance resty-go
	cli := resty.New().
		SetDebug(true).
		SetTimeout(5 * time.Second)

	//repository
	provinceRepo := repos.InitProvinceRepository(false)
	districtRepo := repos.InitDistrictRepository(false)
	wardRepo := repos.InitWardRepository(false)
	deliveryRepo := repos.NewDeliveryRepos(db)

	//service
	userServ := userserv.NewUserService(cli, globalCfg)
	shippingServ := shippingserv.NewShippingCostService(&provinceRepo, &userServ, &deliveryRepo)
	deliveryServ := deliveryserv.NewDeliveryService(&userServ, &deliveryRepo)

	//api handler
	vietnamProvinceApi := handler.NewVietNamProvinceHandle(&provinceRepo, &districtRepo, &wardRepo)
	shippingApi := handler.NewShippingHandle(&shippingServ)
	deliveryApi := handler.NewDeliveryHandle(&deliveryServ)

	//middleware
	authMiddleware := middleware.NewAuthMiddleware(&userServ)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	deli := v1.Group("/delivery")

	delivery := deli.Group("/admin", authMiddleware.RequiredRoles([]string{"ADMIN"}))
	delivery.Get("", deliveryApi.GetAllDeliveries)
	delivery.Get("/:id", deliveryApi.GetDeliveryID)
	delivery.Post("", deliveryApi.CreateDelivery)
	delivery.Patch("/:id", deliveryApi.UpdateDelivery)
	delivery.Patch("/:id/status", deliveryApi.UpdateStatusDelivery)

	local := deli.Group("/vn-location")
	local.Get("/province", vietnamProvinceApi.GetAllProvince)
	local.Get("/district/:id", vietnamProvinceApi.GetAllDistrictByProvince)
	local.Get("/ward/:id", vietnamProvinceApi.GetAllWardByDistrict)

	shipping := deli.Group("/shipping")
	shipping.Post("/anonymous", shippingApi.CalculateShippingByProvinceCode)
	shipping.Post("/order", shippingApi.CalculateOrderShippingCost)

	validate := deli.Group("/validate", authMiddleware.RequiredRoles([]string{"DELIVERY"}))
	validate.Get("", deliveryApi.GetDeliveryByToken)

	err = app.Listen(":5005")
	if err != nil {
		return
	}
}
