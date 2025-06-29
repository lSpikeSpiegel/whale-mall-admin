package service

import (
	"whale/mall/admin/model"
)

func AddToCart(userID, productID uint, quantity int) error {
	var cartItem model.CartItem

	err := model.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error

	if err != nil {
		// exist
		cartItem.Quantity = quantity
		return model.DB.Save(&cartItem).Error
	}

	// not exist
	netItem := model.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	return model.DB.Create(&netItem).Error
}

func GetCartItems(userID uint) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	err := model.DB.Where("user_id = ?", userID).Preload("Product").Order("created_at DESC").Find(&cartItems).Error
	return cartItems, err
}

func UpdateCartItemQuantity(id uint, quantity int) error {
	return model.DB.Model(&model.CartItem{}).Where("id =?", id).Update("quantity", quantity).Error
}

func DeleteCartItem(id uint) error {
	return model.DB.Delete(&model.CartItem{}, id).Error
}
