package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"

	"github.com/gin-gonic/gin"
)

type LaundryController struct {
	c                 *gin.Context
	LaundryRepository *repositories.LaundryRepository
}

func NewLaundryController(LaundryRepository *repositories.LaundryRepository) *LaundryController {
	return &LaundryController{
		LaundryRepository: LaundryRepository,
	}
}

func (controller LaundryController) Laundries() ([]types.Laundry, error) {
	laundries, err := controller.LaundryRepository.FindAllLaundries()
	return laundries, err
}
