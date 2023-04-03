package laundry

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type CreateLaundryHandler struct {
	repository *repositories.LaundryRepository
}

func NewCreateLaundryHandler(repository *repositories.LaundryRepository) *CreateLaundryHandler {
	return &CreateLaundryHandler{repository: repository}
}

func (ch CreateLaundryHandler) Handle(laundry *types.Laundry) (types.Laundry, error) {
	laundryDb, err := ch.repository.Create(laundry)

	return laundryDb, err
}
