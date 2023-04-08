package address

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type DeleteAddressHandler struct {
	repository *repositories.AddressRepository
}

func NewDeleteAddressHandler(repository *repositories.AddressRepository) *DeleteAddressHandler {
	return &DeleteAddressHandler{
		repository: repository,
	}
}

func (ch *DeleteAddressHandler) Handle(address *types.Address) (*types.Address, error) {
	addressDeleted, err := ch.repository.Delete(address)

	return addressDeleted, err
}
