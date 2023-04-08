package address

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetAddressesHandler struct {
	repository repositories.AddressRepository
}

func (ch *GetAddressesHandler) NewGetAddressesHandler(address *types.Address) (*[]types.Address, error) {
	addressesFound, errFind := ch.repository.GetAddresses(address)

	return addressesFound, errFind
}
