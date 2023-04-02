package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/controllers/handlers/laundry"
)

type LaundryController struct {
	GetLaundriesHandler *laundry.GetLaundriesHandler
}

func NewLaundryController(GetLaundriesHandler *laundry.GetLaundriesHandler) *LaundryController {
	return &LaundryController{
		GetLaundriesHandler: GetLaundriesHandler,
	}
}

func (controller LaundryController) Laundries() ([]types.Laundry, error) {
	laundries, err := controller.GetLaundriesHandler.Handle()
	return laundries, err
}
