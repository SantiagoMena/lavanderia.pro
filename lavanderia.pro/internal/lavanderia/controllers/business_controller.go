package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/business"
)

type BusinessController struct {
	GetAllBusinessHandler       *business.GetAllBusinessHandler
	CreateBusinessHandler       *business.CreateBusinessHandler
	DeleteBusinessHandler       *business.DeleteBusinessHandler
	UpdateBusinessHandler       *business.UpdateBusinessHandler
	GetBusinessHandler          *business.GetBusinessHandler
	GetAllBusinessByAuthHandler *business.GetAllBusinessByAuthHandler
}

func NewBusinessController(
	GetAllBusinessHandler *business.GetAllBusinessHandler,
	CreateBusinessHandler *business.CreateBusinessHandler,
	DeleteBusinessHandler *business.DeleteBusinessHandler,
	UpdateBusinessHandler *business.UpdateBusinessHandler,
	GetBusinessHandler *business.GetBusinessHandler,
	GetAllBusinessByAuthHandler *business.GetAllBusinessByAuthHandler,
) *BusinessController {
	return &BusinessController{
		GetAllBusinessHandler:       GetAllBusinessHandler,
		CreateBusinessHandler:       CreateBusinessHandler,
		DeleteBusinessHandler:       DeleteBusinessHandler,
		UpdateBusinessHandler:       UpdateBusinessHandler,
		GetBusinessHandler:          GetBusinessHandler,
		GetAllBusinessByAuthHandler: GetAllBusinessByAuthHandler,
	}
}

func (controller BusinessController) GetAllBusiness() ([]types.Business, error) {
	business, err := controller.GetAllBusinessHandler.Handle()
	return business, err
}

func (controller BusinessController) PostBusiness(business *types.Business) (types.Business, error) {
	// Handle Create Business
	businessDb, err := controller.CreateBusinessHandler.Handle(business)

	if err != nil {
		return types.Business{}, err
	}

	return businessDb, err
}

func (controller BusinessController) DeleteBusiness(business *types.Business) (types.Business, error) {
	// Handle Create Business
	businessDb, err := controller.DeleteBusinessHandler.Handle(business)

	if err != nil {
		return types.Business{}, err
	}

	return businessDb, err
}

func (controller BusinessController) UpdateBusiness(business *types.Business) (types.Business, error) {
	// Handle Create Business
	businessDb, err := controller.UpdateBusinessHandler.Handle(business)

	if err != nil {
		return types.Business{}, err
	}

	return businessDb, err
}

func (controller BusinessController) GetBusiness(business *types.Business) (types.Business, error) {
	// Handle Create Business
	businessDb, err := controller.GetBusinessHandler.Handle(business)

	if err != nil {
		return types.Business{}, err
	}

	return businessDb, err
}

func (controller BusinessController) GetAllBusinessByAuth(authId string) ([]types.Business, error) {
	business, err := controller.GetAllBusinessByAuthHandler.Handle(authId)
	return business, err
}

// func (controller BusinessController) RegisterDelivery(authId string) (types.Delivery, error) {
// 	business, err := controller.RegisterDeliv.Handle(authId)
// 	return business, err
// }
