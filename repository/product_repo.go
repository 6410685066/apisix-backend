package repository

import (
	"apisix-backend/share"
	"apisix-backend/structs"
)

func FindProducts(para structs.ProductGetRequest) ([]structs.ProductsEntity, error) {
	var products []structs.ProductsEntity

	db := share.DB

	if para.Name != "" {
		db = db.Where("name LIKE ?", "%"+para.Name+"%")
	}
	if para.Category != "" {
		db = db.Where("category = ?", para.Category)
	}
	if para.Price != 0 {
		db = db.Where("price >= ?", para.Price)
	}
	if para.Stock != 0 {
		db = db.Where("stock = ?", para.Stock)
	}

	err := db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func CreateProduct(product *structs.ProductPostRequest) error {
	newProduct := structs.ProductsEntity{
		Name:     product.Name,
		Category: product.Category,
		Price:    product.Price,
		Stock:    product.Stock,
		CreateBy: product.CreateBy,
		UpdateBy: product.UpdateBy,
	}

	err := share.DB.Create(&newProduct).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(product *structs.ProductPutRequest) error {
	existingProduct := structs.ProductsEntity{
		ID:       product.ID,
		Name:     product.Name,
		Category: product.Category,
		Price:    product.Price,
		Stock:    product.Stock,
		CreateBy: product.CreateBy,
		UpdateBy: product.UpdateBy,
	}

	err := share.DB.Save(&existingProduct).Error
	if err != nil {
		return err
	}

	return nil
}

func PatchProduct(product *structs.ProductPatchRequest) error {
	existingProduct := structs.ProductsEntity{
		ID: product.ID,
	}

	if err := share.DB.First(&existingProduct).Error; err != nil {
		return err
	}
	isUpdated := false
	if product.Name != nil && *product.Name != "" {
		isUpdated = true
		existingProduct.Name = *product.Name
	}
	if product.Category != nil && *product.Category != "" {
		isUpdated = true
		existingProduct.Category = *product.Category
	}
	if product.Price != nil && *product.Price >= 0 {
		isUpdated = true
		existingProduct.Price = *product.Price
	}
	if product.Stock != nil {
		isUpdated = true
		existingProduct.Stock = *product.Stock
	}
	if isUpdated {
		existingProduct.UpdateBy = product.UpdateBy
	}

	err := share.DB.Save(&existingProduct).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id uint) error {
	var product structs.ProductsEntity
	if err := share.DB.First(&product, id).Error; err != nil {
		return err
	}

	err := share.DB.Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}
