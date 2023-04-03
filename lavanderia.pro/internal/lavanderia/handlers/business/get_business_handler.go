package business

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetBusinessHandler struct {
	repository *repositories.BusinessRepository
}

func NewGetBusinessHandler(repository *repositories.BusinessRepository) *GetBusinessHandler {
	return &GetBusinessHandler{repository: repository}
}

func (ch GetBusinessHandler) Handle(business *types.Business) (types.Business, error) {
	businessDb, err := ch.repository.Get(business)

	return businessDb, err
}
