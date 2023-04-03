package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/laundry"
)

type LaundryController struct {
	GetLaundriesHandler  *laundry.GetLaundriesHandler
	CreateLaundryHandler *laundry.CreateLaundryHandler
}

func NewLaundryController(GetLaundriesHandler *laundry.GetLaundriesHandler, CreateLaundryHandler *laundry.CreateLaundryHandler) *LaundryController {
	return &LaundryController{
		GetLaundriesHandler:  GetLaundriesHandler,
		CreateLaundryHandler: CreateLaundryHandler,
	}
}

func (controller LaundryController) GetLaundries() ([]types.Laundry, error) {
	laundries, err := controller.GetLaundriesHandler.Handle()
	return laundries, err
}

func (controller LaundryController) PostLaundry(laundry *types.Laundry) (types.Laundry, error) {
	// // Handle Create Laundry
	laundryDb, err := controller.CreateLaundryHandler.Handle(laundry)

	if err != nil {
		return types.Laundry{}, err
	}

	return laundryDb, err
}
