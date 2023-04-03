package laundry

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type UpdateLaundryHandler struct {
	repository *repositories.LaundryRepository
}

func NewUpdateLaundryHandler(repository *repositories.LaundryRepository) *UpdateLaundryHandler {
	return &UpdateLaundryHandler{repository: repository}
}

func (ch UpdateLaundryHandler) Handle(laundry *types.Laundry) (types.Laundry, error) {
	laundryDb, err := ch.repository.Update(laundry)

	return laundryDb, err
}
