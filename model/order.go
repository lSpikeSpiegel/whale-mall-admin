package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      uint        `gorm:"not null" json:"user_id"`
	OrderNo     string      `gorm:"unique;not null" json:"order_no"`
	TotalPrice  int         `gorm:"not null" json:"total_price"`
	Status      int         `gorm:"not null" json:"status"`   // unpaid / paid / canceled
	PayType     int         `gorm:"not null" json:"pay_type"` // alipay / wechat / unionpay / balance
	PayAt       *time.Time  `json:"pay_at"`
	Delivered   bool        `gorm:"default:false" json:"delivered"`
	DeliveredAt *time.Time  `json:"delivered_at"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}

type OrderItem struct {
	gorm.Model
	OrderNo    string  `gorm:"not null" json:"order_no"`
	OrderID    uint    `gorm:"not null" json:"order_id"`
	ProductID  uint    `gorm:"not null" json:"product_id"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	TotalPrice int     `gorm:"not null" json:"total_price"`
	Product    Product `gorm:"foreignKey:ProductID" json:"product"`
}
