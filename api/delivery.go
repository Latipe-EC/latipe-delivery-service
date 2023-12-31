package api

import (
	"delivery-service/domain/dto"
	"delivery-service/middleware"
	"delivery-service/pkgs/valid"
	"delivery-service/service/deliveryserv"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
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
		resp := dto.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}
	return ctx.JSON(dataResp.Items)
}

func (receiver DeliveryHandle) CreateDelivery(ctx *fiber.Ctx) error {
	context := ctx.Context()

	reqBody := dto.CreateDeliveryRequest{}

	if err := ctx.BodyParser(&reqBody); err != nil {
		resp := dto.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}

	if err := valid.GetValidator().Validate(reqBody); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	token := fmt.Sprintf("%s", ctx.Locals(middleware.BEARER_TOKEN))
	if token == "" {
		resp := dto.DefaultResponse{
			Success: false,
			Message: "",
		}
		return ctx.Status(http.StatusUnauthorized).JSON(resp)
	}

	reqBody.BearerToken = token

	lastId, err := receiver.service.CreateDelivery(context, &reqBody)
	if err != nil {
		resp := dto.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}
	resp := make(map[string]string)
	resp["id"] = lastId

	return ctx.JSON(resp)
}

func (receiver DeliveryHandle) UpdateDelivery(ctx *fiber.Ctx) error {
	context := ctx.Context()

	reqBody := dto.UpdateDeliveryRequest{}

	if err := ctx.BodyParser(&reqBody); err != nil {
		resp := dto.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}

	if err := ctx.ParamsParser(&reqBody); err != nil {
		resp := dto.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}

	if err := valid.GetValidator().Validate(reqBody); err != nil {
		resp := dto.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusBadRequest).JSON(resp)
	}

	err := receiver.service.UpdateDelivery(context, &reqBody)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	resp := dto.UpdateDeliveryResponse{
		Status:  true,
		Message: "the delivery was updated",
	}

	return ctx.JSON(resp)
}

func (receiver DeliveryHandle) UpdateStatusDelivery(ctx *fiber.Ctx) error {
	context := ctx.Context()

	reqBody := dto.UpdateStatusDeliveryRequest{}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if err := ctx.ParamsParser(&reqBody); err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if err := valid.GetValidator().Validate(reqBody); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err := receiver.service.UpdateStatusDelivery(context, &reqBody)
	if err != nil {
		resp := dto.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := dto.UpdateDeliveryResponse{
		Status:  true,
		Message: "the delivery was updated",
	}

	return ctx.JSON(resp)
}

func (receiver DeliveryHandle) GetDeliveryByToken(ctx *fiber.Ctx) error {
	context := ctx.Context()

	userId := fmt.Sprintf("%s", ctx.Locals(middleware.USER_ID))
	if userId == "" {
		resp := dto.DefaultResponse{
			Success: false,
			Message: "",
		}
		return ctx.Status(http.StatusUnauthorized).JSON(resp)
	}

	resp, err := receiver.service.GetByUserId(context, userId)
	if err != nil {
		resp := dto.DefaultResponse{
			Success: false,
			Message: "",
		}

		if err == mongo.ErrNoDocuments {
			resp.Message = "not found"
			return ctx.Status(http.StatusNotFound).JSON(resp)
		}

		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}

	return ctx.JSON(resp)
}

func (receiver DeliveryHandle) GetDeliveryID(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := dto.GetDeliveryByIdRequest{}

	if err := ctx.ParamsParser(&req); err != nil {
		resp := dto.DefaultResponse{
			Success: false,
			Message: "",
		}
		return ctx.Status(http.StatusBadRequest).JSON(resp)
	}

	resp, err := receiver.service.GetById(context, req.DeliveryId)
	if err != nil {
		resp := dto.DefaultResponse{
			Success: false,
			Message: "",
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}

	return ctx.JSON(resp)
}
