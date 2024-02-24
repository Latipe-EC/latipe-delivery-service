package api

import (
	_ "delivery-service/docs"
	"delivery-service/internal/domain/dto"
	"delivery-service/internal/service/shippingserv"
	"delivery-service/pkgs/valid"
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

// @Summary Calculate shipping cost by province code
// @Tags Shipping
// @Description Calculate shipping cost by province code
// @Accept json
// @Produce json
// @Param body body CalculateShippingCostRequest true "Calculate Shipping Cost Request"
// @Router /shipping/anonymous [post]
// @Success 200 {object} CalculateShippingCostResponse
// @Failure 400 {object} DefaultResponse
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

// @Summary Calculate order shipping cost
// @Tags Shipping
// @Description Calculate order shipping cost
// @Accept json
// @Produce json
// @Param body body OrderShippingCostRequest true "Order Shipping Cost Request"
// @Router /shipping/order [post]
// @Success 200 {object} OrderShippingCostResponse
// @Failure 400 {object} DefaultResponse
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
