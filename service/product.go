package service

import "whale/mall/admin/model"

func CreateProduct(product *model.Product) error {
	return model.DB.Create(product).Error
}

func GetProductList(page int, size int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	db := model.DB.Model(&model.Product{})

	// 获取商品总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * size).Limit(size).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	return products, total, err
}

func GetProductById(id int) (*model.Product, error) {
	var product model.Product
	err := model.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, err
}
