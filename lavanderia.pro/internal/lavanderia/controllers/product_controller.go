package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/product"
)

type ProductController struct {
	CreateProductHandler            *product.CreateProductHandler
	GetAllProductsByBusinessHandler *product.GetAllProductsByBusinessHandler
}

func NewProductController(
	CreateProductHandler *product.CreateProductHandler,
	GetAllProductsByBusinessHandler *product.GetAllProductsByBusinessHandler,
) *ProductController {
	return &ProductController{
		CreateProductHandler:            CreateProductHandler,
		GetAllProductsByBusinessHandler: GetAllProductsByBusinessHandler,
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

func (controller ProductController) GetAllProductsByBusiness(business string) ([]types.Product, error) {
	// Handle Create Product
	productsDb, err := controller.GetAllProductsByBusinessHandler.Handle(business)

	if err != nil {
		return []types.Product{}, err
	}

	return productsDb, err
}
