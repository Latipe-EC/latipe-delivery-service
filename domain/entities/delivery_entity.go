package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Delivery struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DeliveryName string             `json:"delivery_name,omitempty" bson:"delivery_name,omitempty"`
	DeliveryCode string             `json:"delivery_code,omitempty"  bson:"delivery_code,omitempty"`
	BaseCost     int                `json:"base_cost,omitempty" bson:"base_cost,omitempty"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	IsActive     bool               `json:"is_active,omitempty" bson:"is_active,omitempty"`
}
