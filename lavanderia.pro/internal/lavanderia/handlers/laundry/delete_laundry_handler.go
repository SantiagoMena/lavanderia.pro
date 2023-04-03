package laundry

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type DeleteLaundryHandler struct {
	repository *repositories.LaundryRepository
}

func NewDeleteLaundryHandler(repository *repositories.LaundryRepository) *DeleteLaundryHandler {
	return &DeleteLaundryHandler{repository: repository}
}

func (ch DeleteLaundryHandler) Handle(laundry *types.Laundry) (types.Laundry, error) {
	laundryDb, err := ch.repository.Delete(laundry)

	return laundryDb, err
}
