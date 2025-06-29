package router

import (
	"whale/mall/admin/controller"
	"whale/mall/admin/middleware"

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

	auth := api.Group("/")
	auth.Use(middleware.JWTAuth())

	// product
	{
		auth.POST("/product", controller.CreateProduct)
		api.GET("/products", controller.GetProductList)
		api.GET("/product/:id", controller.GetProductDetail)
	}

	// category
	{
		auth.POST("/category", controller.CreateCategory)
		api.GET("/categories", controller.GetCategoryList)
	}

	// cart
	{
		auth.POST("/cart", controller.AddToCart)
		auth.GET("/cart", controller.ListCart)
		auth.PUT("/cart/:id", controller.UpdateCart)
		auth.DELETE("/cart/:id", controller.DeleteCart)
	}

	return r
}
