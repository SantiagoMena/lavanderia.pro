package controllers

import (
	"lavanderia.pro/api/types"

	"lavanderia.pro/internal/lavanderia/repositories"

	"github.com/gin-gonic/gin"
)

func GetLaundries() []types.Laundry {
	return repositories.FindAllLaundries()
}

func Laundries(c *gin.Context) {
	c.JSON(200, repositories.FindAllLaundries())
}
