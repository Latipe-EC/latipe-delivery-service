package dto

type CalculateShippingCostRequest struct {
	SrcCode  []string `json:"src_code" validate:"required"`
	DestCode string   `json:"dest_code" validate:"required"`
}

type OrderShippingCostRequest struct {
	SrcCode    string `json:"src_code" validate:"required"`
	DestCode   string `json:"dest_code" validate:"required"`
	DeliveryId string `json:"delivery_id" validate:"required"`
}

type CalculateShippingCostShipping struct {
	SrcCode      []string `json:"src_code,omitempty"`
	DestCode     string   `json:"dest_code,omitempty"`
	ReceiveDate  string   `json:"receive_date"`
	DeliveryId   string   `json:"delivery_id"`
	DeliveryName string   `json:"delivery_name"`
	Cost         int      `json:"cost"`
}
