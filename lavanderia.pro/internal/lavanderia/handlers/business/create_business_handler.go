package business

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type CreateBusinessHandler struct {
	repository *repositories.BusinessRepository
}

func NewCreateBusinessHandler(repository *repositories.BusinessRepository) *CreateBusinessHandler {
	return &CreateBusinessHandler{repository: repository}
}

func (ch CreateBusinessHandler) Handle(business *types.Business) (types.Business, error) {
	businessDb, err := ch.repository.Create(business)

	return businessDb, err
}
