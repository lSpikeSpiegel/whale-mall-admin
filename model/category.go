package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null;unique" json:"name"`
	ParentId uint   `gorm:"type:int;not null" json:"parent_id"`
}
