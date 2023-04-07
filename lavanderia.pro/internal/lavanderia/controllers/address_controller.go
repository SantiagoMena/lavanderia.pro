package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/address"
)

type AddressController struct {
	CreateAddressHandler *address.CreateAddressHandler
	GetAddressHandler    *address.GetAddressHandler
}

func NewAddressController(
	CreateAddressHandler *address.CreateAddressHandler,
	GetAddressHandler *address.GetAddressHandler,
) *AddressController {
	return &AddressController{
		CreateAddressHandler: CreateAddressHandler,
		GetAddressHandler:    GetAddressHandler,
	}
}

func (controller *AddressController) CreateAddress(address *types.Address) (types.Address, error) {
	addressCreated, err := controller.CreateAddressHandler.Handle(address)

	return addressCreated, err
}

func (controller *AddressController) GetAddress(address *types.Address) (*types.Address, error) {
	addressFound, err := controller.GetAddressHandler.Handle(address)

	return addressFound, err
}
