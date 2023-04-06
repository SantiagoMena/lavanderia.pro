package product

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type DeleteProductHandler struct {
	repository *repositories.ProductRepository
}

func NewDeleteProductHandler(repository *repositories.ProductRepository) *DeleteProductHandler {
	return &DeleteProductHandler{repository: repository}
}

func (ch DeleteProductHandler) Handle(product *types.Product) (types.Product, error) {
	productDb, err := ch.repository.Delete(product)

	return productDb, err
}
