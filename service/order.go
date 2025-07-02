package service

import (
	"whale/mall/admin/model"
	"whale/mall/admin/utils"

	"gorm.io/gorm"
)

func CreateOrder(cartItems []*model.CartItem) (string, error) {
	var totalPrice int
	var orderItems []model.OrderItem

	orderNo, err := utils.GenerateOrderNo()
	if err != nil {
		return "", err
	}

	for _, item := range cartItems {
		price := item.Product.Price
		amount := price * item.Quantity
		totalPrice += amount
		orderItems = append(orderItems, model.OrderItem{
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			TotalPrice: amount,
			OrderNo:    orderNo,
		})
	}

	order := model.Order{
		UserID:     cartItems[0].UserID,
		OrderNo:    orderNo,
		TotalPrice: totalPrice,
		OrderItems: orderItems,
	}

	err2 := model.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		var cartIDs []uint
		for _, item := range cartItems {
			cartIDs = append(cartIDs, item.ID)
		}
		tx.Where("id IN ?", cartIDs).Delete(&model.CartItem{})

		return nil
	})

	if err2 != nil {
		return "", err2
	}

	return orderNo, nil

}
