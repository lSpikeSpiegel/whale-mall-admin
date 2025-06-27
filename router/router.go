package router

import (
	"whale/mall/admin/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	// user
	{
		api.POST("/register", controller.Register)
		api.POST("/login", controller.Login)
	}

	// product
	{
		api.POST("/product", controller.CreateProduct)
		api.GET("/products", controller.GetProductList)
		api.GET("/product/:id", controller.GetProductDetail)
	}

	// category
	{
		api.POST("/category", controller.CreateCategory)
		api.GET("/categories", controller.GetCategoryList)
	}

	return r
}
