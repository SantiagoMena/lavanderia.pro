package address

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type UpdateAddressHandler struct {
	repository *repositories.AddressRepository
}

func NewUpdateAddressHandler(repository *repositories.AddressRepository) *UpdateAddressHandler {
	return &UpdateAddressHandler{
		repository: repository,
	}
}

func (ch *UpdateAddressHandler) Handle(address *types.Address) (*types.Address, error) {
	addressUpdated, err := ch.repository.Update(address)

	return addressUpdated, err
}
