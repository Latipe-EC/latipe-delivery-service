package api

import (
	"delivery-service/domain/dto"
	"delivery-service/service/shippingserv"
	"delivery-service/valid"
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

	if err := valid.GetValidator().Validate(request); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	dataResp, err := api.service.CalculateByProvinceCode(ctx.Context(), &request)
	if err != nil {
		log.Errorf("%v", err)
		return ctx.Status(http.StatusBadRequest).SendString("Invalid params")
	}

	return ctx.JSON(dataResp)

}

func (api ShippingHandle) CalculateOrderShippingCost(ctx *fiber.Ctx) error {
	var request dto.OrderShippingCostRequest

	if err := ctx.BodyParser(&request); err != nil {
		log.Errorf("%v", err)
		return ctx.Status(http.StatusBadRequest).SendString("Parse body was failed")
	}

	if err := valid.GetValidator().Validate(request); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	dataResp, err := api.service.CalculateOrderShippingCost(ctx.Context(), &request)
	if err != nil {
		log.Errorf("%v", err)
		return ctx.Status(http.StatusBadRequest).SendString("Invalid params")
	}

	return ctx.JSON(dataResp)

}
