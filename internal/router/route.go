package router

import (
	"delivery-service/internal/api"
	"delivery-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type RouterHandler struct {
	deliveryApi        *api.DeliveryHandle
	shippingApi        *api.ShippingHandle
	vietnamProvinceApi *api.VietNamProvinceHandle
	authMiddleware     *middleware.AuthMiddleware
}

func NewRouterHandler(deliveryApi *api.DeliveryHandle,
	shippingApi *api.ShippingHandle,
	vietnamProvinceApi *api.VietNamProvinceHandle,
	authMiddleware *middleware.AuthMiddleware) *RouterHandler {
	return &RouterHandler{
		deliveryApi:        deliveryApi,
		shippingApi:        shippingApi,
		vietnamProvinceApi: vietnamProvinceApi,
		authMiddleware:     authMiddleware,
	}
}

func (rt RouterHandler) InitRouter(root *fiber.Router) {

	deli := (*root).Group("/delivery")

	delivery := deli.Group("/admin", rt.authMiddleware.RequiredRoles([]string{"ADMIN"}))
	delivery.Get("", rt.deliveryApi.GetAllDeliveries)
	delivery.Get("/:id", rt.deliveryApi.GetDeliveryID)
	delivery.Post("", rt.deliveryApi.CreateDelivery)
	delivery.Patch("/:id", rt.deliveryApi.UpdateDelivery)
	delivery.Patch("/:id/status", rt.deliveryApi.UpdateStatusDelivery)

	local := deli.Group("/vn-location")
	local.Get("/province", rt.vietnamProvinceApi.GetAllProvince)
	local.Get("/district/:id", rt.vietnamProvinceApi.GetAllDistrictByProvince)
	local.Get("/ward/:id", rt.vietnamProvinceApi.GetAllWardByDistrict)

	shipping := deli.Group("/shipping")
	shipping.Post("/anonymous", rt.shippingApi.CalculateShippingByProvinceCode)
	shipping.Post("/order", rt.shippingApi.CalculateOrderShippingCost)

	validate := deli.Group("/validate", rt.authMiddleware.RequiredRoles([]string{"DELIVERY"}))
	validate.Get("", rt.deliveryApi.GetDeliveryByToken)
}
