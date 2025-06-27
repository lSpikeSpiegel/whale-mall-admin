package model

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	UserID    uint `gorm:"not null" json:"user_id"`
	ProductID uint `gorm:"not null" json:"product_id"`
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}
