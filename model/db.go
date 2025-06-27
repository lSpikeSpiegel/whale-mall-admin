package model

import (
	"fmt"
	"whale/mall/admin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var models = []interface{}{
	&User{},
	&Product{},
	&Category{},
	&CartItem{},
	// 后续在这里注册新模型
}

func InitDB() {
	d := config.GlobalConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.Port, d.DbName, d.Charset)

	fmt.Println(dsn)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	DB.AutoMigrate(models...)

}
