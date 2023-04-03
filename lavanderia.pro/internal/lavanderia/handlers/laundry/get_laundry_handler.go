package laundry

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetLaundryHandler struct {
	repository *repositories.LaundryRepository
}

func NewGetLaundryHandler(repository *repositories.LaundryRepository) *GetLaundryHandler {
	return &GetLaundryHandler{repository: repository}
}

func (ch GetLaundryHandler) Handle(laundry *types.Laundry) (types.Laundry, error) {
	laundryDb, err := ch.repository.Get(laundry)

	return laundryDb, err
}
