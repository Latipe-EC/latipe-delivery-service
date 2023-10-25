package dto

type CreateDeliveryRequest struct {
	DeliveryName string `json:"delivery_name" validate:"required"`
	DeliveryCode string `json:"delivery_code" validate:"required"`
	BaseCost     int    `json:"base_cost" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type CreateDeliveryResponse struct {
	ID string `json:"_id" `
}
