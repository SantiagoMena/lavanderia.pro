package business

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type SearchBusinessHandler struct {
	repository *repositories.BusinessRepository
}

func NewSearchBusinessHandler(repository *repositories.BusinessRepository) *SearchBusinessHandler {
	return &SearchBusinessHandler{repository: repository}
}

func (ch SearchBusinessHandler) Handle() ([]types.Business, error) {
	businessDb, err := ch.repository.Search()

	return businessDb, err
}
