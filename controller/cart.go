package controller

import (
	"strconv"
	"whale/mall/admin/service"

	"github.com/gin-gonic/gin"
)

type AddCartRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func AddToCart(c *gin.Context) {
	userId, exist := c.Get("userId")

	if !exist {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var req AddCartRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Quantity <= 0 {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}

	err := service.AddToCart(userId.(uint), req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func ListCart(c *gin.Context) {
	userId, exist := c.Get("userId")
	if !exist {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	cart, err := service.GetCartItems(userId.(uint))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": cart})
}

func UpdateCart(c *gin.Context) {
	_, exist := c.Get("userId")
	if !exist {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req AddCartRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Quantity <= 0 {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}

	err := service.UpdateCartItemQuantity(uint(id), req.Quantity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func DeleteCart(c *gin.Context) {
	_, exist := c.Get("userId")
	if !exist {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := service.DeleteCartItem(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
