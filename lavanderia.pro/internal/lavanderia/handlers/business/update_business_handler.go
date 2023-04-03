package business

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type UpdateBusinessHandler struct {
	repository *repositories.BusinessRepository
}

func NewUpdateBusinessHandler(repository *repositories.BusinessRepository) *UpdateBusinessHandler {
	return &UpdateBusinessHandler{repository: repository}
}

func (ch UpdateBusinessHandler) Handle(business *types.Business) (types.Business, error) {
	businessDb, err := ch.repository.Update(business)

	return businessDb, err
}
