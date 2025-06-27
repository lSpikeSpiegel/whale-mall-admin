package controller

import (
	"strconv"
	"whale/mall/admin/model"
	"whale/mall/admin/service"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}

	if err := service.CreateProduct(&product); err != nil {
		c.JSON(500, gin.H{"error": "failed to create product"})
		return
	}
	c.JSON(200, gin.H{"message": "product created successfully"})
}

func GetProductList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	products, total, err := service.GetProductList(page, size)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get product list"})
		return
	}
	c.JSON(200, gin.H{"data": products, "total": total})
}

func GetProductDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}

	product, err := service.GetProductById(id)

	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get product detail"})
		return
	}
	c.JSON(200, gin.H{"data": product})
}
