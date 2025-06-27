package controller

import (
	"fmt"
	"whale/mall/admin/model"
	"whale/mall/admin/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req struct {
		Name     string
		Password string
		Mobile   string
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}

	hashed, err := utils.HashPassword(req.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": "failed to hash password"})
		return
	}

	user := model.User{
		Name:     req.Name,
		Password: hashed,
		Mobile:   req.Mobile,
	}

	if err := model.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "user created successfully"})
}

func Login(c *gin.Context) {
	var req struct {
		Name     string
		Mobile   string
		Password string
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}

	var user model.User
	if err := model.DB.Where("name = ? OR mobile = ?", req.Name, req.Mobile).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "invalid username or password"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(400, gin.H{"error": "invalid username or password"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)

	c.JSON(200, gin.H{"token": token})
}
