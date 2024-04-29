package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ShippingPackage struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrderID       string             `json:"order_id" bson:"order_id"`
	DeliveryID    string             `json:"delivery_id" bson:"delivery_id"`
	DeliveryName  string             `json:"delivery_name" bson:"delivery_name"`
	Address       Address            `json:"address" bson:"address"`
	Status        int                `json:"status"bson:"status"`
	PaymentMethod int                `json:"payment_method" bson:"payment_method"`
	Total         int                `json:"total" bson:"total"`
	ShippingCost  int                `json:"shipping_cost" bson:"shipping_cost"`
	ReceiveDate   time.Time          `json:"receive_date" bson:"receive_date"`
	ShippingItems []ShippingItems    `json:"items" bson:"shipping_items"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" json:"updated_at	"`
}

type Address struct {
	Name          string `json:"name" bson:"name"`
	Phone         string `json:"phone" bson:"phone"`
	AddressDetail string `json:"address_detail" bson:"address_detail"`
	ProvinceCode  string `json:"province_code" bson:"province_code"`
	DistrictCode  string `json:"district_code" bson:"district_code"`
	WardCode      string `json:"ward_code" bson:"ward_code"`
}

type ShippingItems struct {
	ProductName  string `json:"product_name" bson:"product_name"`
	ProductID    string `json:"product_id" bson:"product_id"`
	ProductImage string `json:"product_image" bson:"product_image"`
	Quantity     int    `json:"quantity" bson:"quantity"`
}
