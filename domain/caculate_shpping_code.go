package domain

type CalculateShippingCostRequest struct {
	SrcCode    string `json:"src_code"`
	DestCode   string `json:"dest_code"`
	DeliveryId string `json:"delivery_id"`
}

type CalculateShippingCostShipping struct {
	SrcCode      string `json:"src_code"`
	DestCode     string `json:"dest_code"`
	ReceiveDate  string `json:"receive_date"`
	DeliveryId   string `json:"delivery_id"`
	DeliveryName string `json:"delivery_name"`
	Cost         int    `json:"cost"`
}
