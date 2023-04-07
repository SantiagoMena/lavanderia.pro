package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/address"
)

type AddressController struct {
	CreateAddressHandler *address.CreateAddressHandler
}

func NewAddressController(
	CreateAddressHandler *address.CreateAddressHandler,
) *AddressController {
	return &AddressController{
		CreateAddressHandler: CreateAddressHandler,
	}
}

func (controller *AddressController) CreateAddress(address *types.Address) (types.Address, error) {
	addressCreated, err := controller.CreateAddressHandler.Handle(address)

	return addressCreated, err
}
