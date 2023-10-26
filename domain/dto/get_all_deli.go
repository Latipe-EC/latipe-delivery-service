package dto

import (
	"time"
)

type GetAllDeliveries struct {
	Items []DeliveryDetail
}

type DeliveryDetail struct {
	ID           string    `json:"_id" `
	DeliveryName string    `json:"delivery_name,omitempty"`
	DeliveryCode string    `json:"delivery_code,omitempty"`
	BaseCost     int       `json:"base_cost,omitempty" `
	Description  string    `json:"description,omitempty" `
	CreateAt     time.Time `json:"create_at,omitempty"`
	UpdateAt     time.Time `json:"update_at,omitempty"`
	IsActive     bool      `json:"is_active"`
}
