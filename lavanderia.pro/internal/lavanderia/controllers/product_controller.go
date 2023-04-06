package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/product"
)

type ProductController struct {
	CreateProductHandler            *product.CreateProductHandler
	GetAllProductsByBusinessHandler *product.GetAllProductsByBusinessHandler
	DeleteProductHandler            *product.DeleteProductHandler
	GetProductHandler               *product.GetProductHandler
	UpdateProductHandler            *product.UpdateProductHandler
}

func NewProductController(
	CreateProductHandler *product.CreateProductHandler,
	GetAllProductsByBusinessHandler *product.GetAllProductsByBusinessHandler,
	DeleteProductHandler *product.DeleteProductHandler,
	GetProductHandler *product.GetProductHandler,
	UpdateProductHandler *product.UpdateProductHandler,
) *ProductController {
	return &ProductController{
		CreateProductHandler:            CreateProductHandler,
		GetAllProductsByBusinessHandler: GetAllProductsByBusinessHandler,
		DeleteProductHandler:            DeleteProductHandler,
		GetProductHandler:               GetProductHandler,
		UpdateProductHandler:            UpdateProductHandler,
	}
}

func (controller ProductController) PostProduct(product *types.Product) (types.Product, error) {
	productDb, err := controller.CreateProductHandler.Handle(product)

	if err != nil {
		return types.Product{}, err
	}

	return productDb, err
}

func (controller ProductController) GetAllProductsByBusiness(business string) ([]types.Product, error) {
	productsDb, err := controller.GetAllProductsByBusinessHandler.Handle(business)

	if err != nil {
		return []types.Product{}, err
	}

	return productsDb, err
}

func (controller ProductController) DeleteProduct(product *types.Product) (types.Product, error) {
	productDb, err := controller.DeleteProductHandler.Handle(product)

	if err != nil {
		return types.Product{}, err
	}

	return productDb, err
}

func (controller ProductController) GetProduct(product *types.Product) (types.Product, error) {
	productDb, err := controller.GetProductHandler.Handle(product)

	if err != nil {
		return types.Product{}, err
	}

	return productDb, err
}

func (controller ProductController) UpdateProduct(product *types.Product) (types.Product, error) {
	productDb, err := controller.UpdateProductHandler.Handle(product)

	if err != nil {
		return types.Product{}, err
	}

	return productDb, err
}
