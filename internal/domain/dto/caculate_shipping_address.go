package dto

type CalculateShippingAddressCostRequest struct {
	StoreAddId string `json:"store_add_id"`
	UserAddId  string `json:"user_add_id"`
	DeliveryId string `json:"delivery_id"`
}

type CalculateShippingAddressCostShipping struct {
	StoreAddId   string `json:"store_add_id"`
	UserAddId    string `json:"user_add_id"`
	ReceiveDate  string `json:"receive_date"`
	DeliveryId   string `json:"delivery_id"`
	DeliveryName string `json:"delivery_name"`
	Cost         int    `json:"cost"`
}
