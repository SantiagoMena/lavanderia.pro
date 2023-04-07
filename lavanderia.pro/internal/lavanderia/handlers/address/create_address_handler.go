package address

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type CreateAddressHandler struct {
	repository *repositories.AddressRepository
}

func NewCreateAddressHandler(repository *repositories.AddressRepository) *CreateAddressHandler {
	return &CreateAddressHandler{
		repository: repository,
	}
}

func (ch *CreateAddressHandler) Handle(address *types.Address) (types.Address, error) {
	newAddress, err := ch.repository.Create(address)

	return newAddress, err
}
