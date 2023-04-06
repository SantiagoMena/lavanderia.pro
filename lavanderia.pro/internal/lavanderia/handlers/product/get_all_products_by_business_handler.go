package product

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetAllProductsByBusinessHandler struct {
	repository *repositories.ProductRepository
}

func NewGetAllProductsByBusinessHandler(repository *repositories.ProductRepository) *GetAllProductsByBusinessHandler {
	return &GetAllProductsByBusinessHandler{repository: repository}
}

func (ch GetAllProductsByBusinessHandler) Handle(business string) ([]types.Product, error) {
	productDb, err := ch.repository.GetAllProductsByBusiness(business)

	return productDb, err
}
