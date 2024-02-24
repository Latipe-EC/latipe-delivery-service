package api

import (
	_ "delivery-service/docs"
	dto2 "delivery-service/internal/domain/dto"
	"delivery-service/internal/middleware"
	"delivery-service/internal/service/deliveryserv"
	"delivery-service/pkgs/valid"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// @title API Documentation
// @version 1.0.1
// @description This is a server for Latipe Delivery Service.
// @host localhost:5005
// @BasePath /api/v2/delivery
// @schemes http
type DeliveryHandle struct {
	service *deliveryserv.DeliveryService
}

func NewDeliveryHandle(service *deliveryserv.DeliveryService) *DeliveryHandle {
	return &DeliveryHandle{
		service: service,
	}
}

// @Summary Get all deliveries
// @Tags Delivery
// @Description Get all deliveries
// @Accept json
// @Produce json
// @Security Bearer
// @Router /admin [get]
func (receiver DeliveryHandle) GetAllDeliveries(ctx *fiber.Ctx) error {
	context := ctx.Context()
	dataResp, err := receiver.service.GetAllDeliveries(context)
	if err != nil {
		resp := dto2.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}
	return ctx.JSON(dataResp.Items)
}

// @Summary Create a delivery
// @Tags Delivery
// @Description Create a delivery
// @Accept json
// @Produce json
// @Security Bearer
// @Router /admin [post]
func (receiver DeliveryHandle) CreateDelivery(ctx *fiber.Ctx) error {
	context := ctx.Context()

	reqBody := dto2.CreateDeliveryRequest{}

	if err := ctx.BodyParser(&reqBody); err != nil {
		resp := dto2.DefaultResponse{
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
		resp := dto2.DefaultResponse{
			Success: false,
			Message: "",
		}
		return ctx.Status(http.StatusUnauthorized).JSON(resp)
	}

	reqBody.BearerToken = token

	lastId, err := receiver.service.CreateDelivery(context, &reqBody)
	if err != nil {
		resp := dto2.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}
	resp := make(map[string]string)
	resp["id"] = lastId

	return ctx.JSON(resp)
}

// @Summary Update a delivery
// @Tags Delivery
// @Description Update a delivery
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Delivery ID"
// @Router /admin/{id} [patch]
func (receiver DeliveryHandle) UpdateDelivery(ctx *fiber.Ctx) error {
	context := ctx.Context()

	reqBody := dto2.UpdateDeliveryRequest{}

	if err := ctx.BodyParser(&reqBody); err != nil {
		resp := dto2.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}

	if err := ctx.ParamsParser(&reqBody); err != nil {
		resp := dto2.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}

	if err := valid.GetValidator().Validate(reqBody); err != nil {
		resp := dto2.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusBadRequest).JSON(resp)
	}

	err := receiver.service.UpdateDelivery(context, &reqBody)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	resp := dto2.UpdateDeliveryResponse{
		Status:  true,
		Message: "the delivery was updated",
	}

	return ctx.JSON(resp)
}

// @Summary Update status a delivery
// @Tags Delivery
// @Description Update status a delivery
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Delivery ID"
// @Router /admin/{id}/status [patch]
func (receiver DeliveryHandle) UpdateStatusDelivery(ctx *fiber.Ctx) error {
	context := ctx.Context()

	reqBody := dto2.UpdateStatusDeliveryRequest{}

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
		resp := dto2.DefaultResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}

	resp := dto2.UpdateDeliveryResponse{
		Status:  true,
		Message: "the delivery was updated",
	}

	return ctx.JSON(resp)
}

// @Summary Get delivery by token
// @Tags Delivery
// @Description Get delivery by token
// @Accept json
// @Produce json
// @Security Bearer
// @Router /admin [get]
func (receiver DeliveryHandle) GetDeliveryByToken(ctx *fiber.Ctx) error {
	context := ctx.Context()

	userId := fmt.Sprintf("%s", ctx.Locals(middleware.USER_ID))
	if userId == "" {
		resp := dto2.DefaultResponse{
			Success: false,
			Message: "",
		}
		return ctx.Status(http.StatusUnauthorized).JSON(resp)
	}

	resp, err := receiver.service.GetByUserId(context, userId)
	if err != nil {
		resp := dto2.DefaultResponse{
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

// @Summary Get delivery by id
// @Tags Delivery
// @Description Get delivery by id
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Delivery ID"
// @Router admin/{id} [get]
func (receiver DeliveryHandle) GetDeliveryID(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := dto2.GetDeliveryByIdRequest{}

	if err := ctx.ParamsParser(&req); err != nil {
		resp := dto2.DefaultResponse{
			Success: false,
			Message: "",
		}
		return ctx.Status(http.StatusBadRequest).JSON(resp)
	}

	resp, err := receiver.service.GetById(context, req.DeliveryId)
	if err != nil {
		resp := dto2.DefaultResponse{
			Success: false,
			Message: "",
		}
		return ctx.Status(http.StatusInternalServerError).JSON(resp)
	}

	return ctx.JSON(resp)
}
