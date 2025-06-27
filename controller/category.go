package controller

import (
	"whale/mall/admin/model"
	"whale/mall/admin/service"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var cat model.Category

	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}

	if err := service.CreateCategory(&cat); err != nil {
		c.JSON(500, gin.H{"error": "failed to create category"})
		return
	}

	c.JSON(200, gin.H{"message": "success", "data": cat})
}

func GetCategoryList(c *gin.Context) {
	cats, err := service.GetAllCategories()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get category list"})
		return
	}
	c.JSON(200, gin.H{"message": "success", "data": cats})
}
