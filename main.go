package main

import (
	"whale/mall/admin/config"
	"whale/mall/admin/model"
	"whale/mall/admin/router"
)

func main() {
	config.InitConfig()
	model.InitDB()

	r := router.InitRouter()
	r.Run(":" + config.GlobalConfig.Server.Port)
}
