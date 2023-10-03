package main

import (
	handler "delivery-service/api"
	"delivery-service/repos"
	"github.com/gofiber/fiber/v2"
)

func main() {
	provinceRepo := repos.InitProvinceRepository()
	districtRepo := repos.InitDistrictRepository()
	wardRepo := repos.InitWardRepository()

	vietnamProvinceApi := handler.NewVietNamProvinceHandle(&provinceRepo, &districtRepo, &wardRepo)

	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	deli := v1.Group("/delivery")
	deli.Get("/province", vietnamProvinceApi.GetAllProvince)
	deli.Get("/district/:id", vietnamProvinceApi.GetAllDistrictByProvince)
	deli.Get("/ward/:id", vietnamProvinceApi.GetAllWardByDistrict)

	err := app.Listen(":5000")
	if err != nil {
		return
	}
}
