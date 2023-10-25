package api

import (
	"delivery-service/domain/dto"
	"delivery-service/service/deliveryserv"
	"delivery-service/valid"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type DeliveryHandle struct {
	service *deliveryserv.DeliveryService
}

func NewDeliveryHandle(service *deliveryserv.DeliveryService) *DeliveryHandle {
	return &DeliveryHandle{
		service: service,
	}
}

func (receiver DeliveryHandle) GetAllDeliveries(ctx *fiber.Ctx) error {
	context := ctx.Context()
	dataResp, err := receiver.service.GetAllDeliveries(context)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.JSON(dataResp.Items)
}

func (receiver DeliveryHandle) CreateDelivery(ctx *fiber.Ctx) error {
	context := ctx.Context()

	reqBody := dto.CreateDeliveryRequest{}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if err := valid.GetValidator().Validate(reqBody); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	lastId, err := receiver.service.CreateDelivery(context, &reqBody)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	resp := make(map[string]string)
	resp["id"] = lastId

	return ctx.JSON(resp)
}
