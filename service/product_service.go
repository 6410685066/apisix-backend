package service

import (
	"apisix-backend/repository"
	"apisix-backend/structs"

	"errors"
)

func GetProductsService(para structs.ProductGetRequest) ([]structs.ProductDto, error) {
	products, err := repository.FindProducts(para)
	if err != nil {
		return nil, err
	}
	productsDto := make([]structs.ProductDto, len(products))
	for i, product := range products {
		productsDto[i] = structs.ProductDto{
			Name:      product.Name,
			Category:  product.Category,
			Price:     product.Price,
			Stock:     product.Stock,
			UpdatedAt: product.UpdatedAt,
		}
	}
	return productsDto, nil
}

func CreateProductService(product *structs.ProductPostRequest) error {
	err := repository.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProductService(id uint, product *structs.ProductPutRequest) error {
	if id != product.ID {
		return errors.New("path ID and body ID do not match")
	}
	err := repository.UpdateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func PatchProductService(id uint, product *structs.ProductPatchRequest) error {
	if id != product.ID {
		return errors.New("path ID and body ID do not match")
	}
	err := repository.PatchProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProductService(id uint) error {
	err := repository.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
