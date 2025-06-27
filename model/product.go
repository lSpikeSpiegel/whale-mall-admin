package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);not null" json:"name"`
	Description  string `gorm:"type:varchar(100)" json:"description"`
	Price        int    `gorm:"type:int;not null" json:"price"`
	Stock        int    `gorm:"type:int" json:"stock"`
	CategoryID   uint   `gorm:"type:int;not null" json:"category_id"`
	ProductImage string `gorm:"type:varchar(255)" json:"product_image"`
}
