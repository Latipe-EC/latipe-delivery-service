package main

import (
	"context"
	"delivery-service/adapter/userserv"
	handler "delivery-service/api"
	"delivery-service/domain/repos"
	"delivery-service/service/shippingserv"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func main() {
	//read env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	//create connect to mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	db := client.Database("delivery_be")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	//create instance fiber
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	//create instance resty-go
	cli := resty.New().
		SetDebug(true).
		SetTimeout(5 * time.Second)

	//repository
	provinceRepo := repos.InitProvinceRepository()
	districtRepo := repos.InitDistrictRepository()
	wardRepo := repos.InitWardRepository()
	deliveryRepo := repos.NewDeliveryRepos(db)

	//service
	userServ := userserv.NewUserService(cli)
	shippingServ := shippingserv.NewShippingCostService(&provinceRepo, &userServ, &deliveryRepo)

	//api handler
	vietnamProvinceApi := handler.NewVietNamProvinceHandle(&provinceRepo, &districtRepo, &wardRepo)
	shippingApi := handler.NewShippingHandle(&shippingServ)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	local := v1.Group("/vn-local")
	local.Get("/province", vietnamProvinceApi.GetAllProvince)
	local.Get("/district/:id", vietnamProvinceApi.GetAllDistrictByProvince)
	local.Get("/ward/:id", vietnamProvinceApi.GetAllWardByDistrict)

	shipping := v1.Group("/shipping")
	shipping.Post("/calculate-cost-anonymous", shippingApi.CalculateShippingByProvinceCode)

	err = app.Listen(":5000")
	if err != nil {
		return
	}
}
