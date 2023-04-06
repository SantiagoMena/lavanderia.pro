package product

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type UpdateProductHandler struct {
	repository *repositories.ProductRepository
}

func NewUpdateProductHandler(repository *repositories.ProductRepository) *UpdateProductHandler {
	return &UpdateProductHandler{repository: repository}
}

func (ch UpdateProductHandler) Handle(product *types.Product) (types.Product, error) {
	productDb, err := ch.repository.Update(product)

	return productDb, err
}
