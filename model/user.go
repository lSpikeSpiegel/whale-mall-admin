package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null;unique" json:"name"`
	Password string `gorm:"type:varchar(225);not null" json:"password"`
	Mobile   string `gorm:"type:varchar(11);not null;unique" json:"mobile"`
}
