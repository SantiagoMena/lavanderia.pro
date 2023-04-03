package business

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type DeleteBusinessHandler struct {
	repository *repositories.BusinessRepository
}

func NewDeleteBusinessHandler(repository *repositories.BusinessRepository) *DeleteBusinessHandler {
	return &DeleteBusinessHandler{repository: repository}
}

func (ch DeleteBusinessHandler) Handle(business *types.Business) (types.Business, error) {
	businessDb, err := ch.repository.Delete(business)

	return businessDb, err
}
