package laundry

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetLaundriesHandler struct {
	repository *repositories.LaundryRepository
}

func NewGetLaundriesHandler(repository *repositories.LaundryRepository) *GetLaundriesHandler {
	return &GetLaundriesHandler{repository: repository}
}

func (ch GetLaundriesHandler) Handle() ([]types.Laundry, error) {
	laundries, err := ch.repository.FindAllLaundries()

	return laundries, err
}
