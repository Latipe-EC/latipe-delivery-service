package api

import (
	"delivery-service/domain/dto"
	"delivery-service/service/shippingserv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"net/http"
)

type ShippingHandle struct {
	service *shippingserv.ShippingCostService
}

func NewShippingHandle(service *shippingserv.ShippingCostService) *ShippingHandle {
	return &ShippingHandle{
		service: service,
	}
}

func (api ShippingHandle) CalculateShippingByProvinceCode(ctx *fiber.Ctx) error {
	var request dto.CalculateShippingCostRequest

	if err := ctx.BodyParser(&request); err != nil {
		log.Errorf("%v", err)
		return ctx.Status(http.StatusBadRequest).SendString("Parse body was failed")
	}

	dataResp, err := api.service.CalculateByProvinceCode(ctx.Context(), &request)
	if err != nil {
		log.Errorf("%v", err)
		return ctx.Status(http.StatusBadRequest).SendString("Invalid params")
	}

	return ctx.JSON(dataResp)

}
