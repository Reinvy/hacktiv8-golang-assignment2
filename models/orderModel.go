package models

import (
	"time"
)

type Order struct {
	OrderId      uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `gorm:"autoCreateTime" json:"ordered_at"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}

type OrderInput struct {
	CustomerName string `json:"customer_name" valid:"required"`
	ItemCode     string `json:"item_code" valid:"required"`
	Description  string `json:"description" valid:"required"`
	Quantity     int    `json:"quantity" valid:"required"`
}
