package product

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type CreateProductHandler struct {
	repository *repositories.ProductRepository
}

func NewCreateProductHandler(repository *repositories.ProductRepository) *CreateProductHandler {
	return &CreateProductHandler{repository: repository}
}

func (ch CreateProductHandler) Handle(product *types.Product) (types.Product, error) {

	productDb, err := ch.repository.Create(product)

	return productDb, err
}
