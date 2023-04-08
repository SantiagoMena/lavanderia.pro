package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/address"
)

type AddressController struct {
	CreateAddressHandler *address.CreateAddressHandler
	GetAddressHandler    *address.GetAddressHandler
	UpdateAddressHandler *address.UpdateAddressHandler
	GetAddressesHandler  *address.GetAddressesHandler
}

func NewAddressController(
	CreateAddressHandler *address.CreateAddressHandler,
	GetAddressHandler *address.GetAddressHandler,
	UpdateAddressHandler *address.UpdateAddressHandler,
	GetAddressesHandler *address.GetAddressesHandler,
) *AddressController {
	return &AddressController{
		CreateAddressHandler: CreateAddressHandler,
		GetAddressHandler:    GetAddressHandler,
		UpdateAddressHandler: UpdateAddressHandler,
		GetAddressesHandler:  GetAddressesHandler,
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

func (controller *AddressController) UpdateAddress(address *types.Address) (*types.Address, error) {
	addressUpdated, err := controller.UpdateAddressHandler.Handle(address)

	return addressUpdated, err
}

func (controller *AddressController) GetAddresses(address *types.Address) (*[]types.Address, error) {
	addressesFound, err := controller.GetAddressesHandler.Handle(address)

	return addressesFound, err
}
