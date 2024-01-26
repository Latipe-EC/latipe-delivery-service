package dto

type CreateDeliveryRequest struct {
	AuthorizationHeader
	DeliveryName string `json:"delivery_name" validate:"required"`
	DeliveryCode string `json:"delivery_code" validate:"required"`
	BaseCost     int    `json:"base_cost" validate:"required"`
	Description  string `json:"description" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
}

type CreateDeliveryResponse struct {
	ID string `json:"_id" `
}
