package business

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetAllBusinessHandler struct {
	repository *repositories.BusinessRepository
}

func NewGetAllBusinessHandler(repository *repositories.BusinessRepository) *GetAllBusinessHandler {
	return &GetAllBusinessHandler{repository: repository}
}

func (ch GetAllBusinessHandler) Handle() ([]types.Business, error) {
	business, err := ch.repository.FindAllBusiness()

	return business, err
}
