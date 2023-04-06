package product

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetProductHandler struct {
	repository *repositories.ProductRepository
}

func NewGetProductHandler(repository *repositories.ProductRepository) *GetProductHandler {
	return &GetProductHandler{repository: repository}
}

func (ch GetProductHandler) Handle(product *types.Product) (types.Product, error) {
	productDb, err := ch.repository.Get(product)

	return productDb, err
}
