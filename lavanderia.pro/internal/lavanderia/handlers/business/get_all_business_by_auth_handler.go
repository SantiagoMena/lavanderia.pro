package business

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetAllBusinessByAuthHandler struct {
	repository *repositories.BusinessRepository
}

func NewGetAllBusinessByAuthHandler(repository *repositories.BusinessRepository) *GetAllBusinessByAuthHandler {
	return &GetAllBusinessByAuthHandler{repository: repository}
}

func (ch GetAllBusinessByAuthHandler) Handle(auth string) ([]types.Business, error) {

	business, err := ch.repository.FindAllBusinessByAuth(auth)

	return business, err
}
