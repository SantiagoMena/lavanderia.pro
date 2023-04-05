package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/product"
)

type ProductController struct {
	CreateProductHandler *product.CreateProductHandler
}

func NewProductController(
	CreateProductHandler *product.CreateProductHandler,
) *ProductController {
	return &ProductController{
		CreateProductHandler: CreateProductHandler,
	}
}

func (controller ProductController) PostProduct(product *types.Product) (types.Product, error) {
	// Handle Create Product
	productDb, err := controller.CreateProductHandler.Handle(product)

	if err != nil {
		return types.Product{}, err
	}

	return productDb, err
}
