package dto

type UpdateDeliveryRequest struct {
	DeliveryName string `json:"delivery_name" validate:"required"`
	DeliveryCode string `json:"delivery_code" validate:"required"`
	BaseCost     int    `json:"base_cost" validate:"required"`
	Description  string `json:"description" validate:"required"`
	DeliId       string `params:"id"`
}

type UpdateStatusDeliveryRequest struct {
	Status bool   `json:"status"`
	DeliId string `params:"id"`
}

type UpdateDeliveryResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
