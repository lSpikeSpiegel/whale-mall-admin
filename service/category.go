package service

import "whale/mall/admin/model"

func CreateCategory(cat *model.Category) (err error) {
	return model.DB.Create(&cat).Error
}

func GetAllCategories() ([]model.Category, error) {
	var list []model.Category
	err := model.DB.Order("created_at asc").Find(&list).Error
	return list, err
}
