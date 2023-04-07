package address

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetAddressHandler struct {
	repository *repositories.AddressRepository
}

func NewGetAddressHandler(repository *repositories.AddressRepository) *GetAddressHandler {
	return &GetAddressHandler{repository: repository}
}

func (ch GetAddressHandler) Handle(address *types.Address) (*types.Address, error) {
	addressDb, err := ch.repository.Get(address)

	return addressDb, err
}
