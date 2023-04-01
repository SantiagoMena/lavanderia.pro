package controllers

import (
	"lavanderia.pro/internal/lavanderia/repositories"

	"github.com/gin-gonic/gin"
)

func Laundries(c *gin.Context) {
	laundries, err := repositories.FindAllLaundries()

	if err != nil {
		c.JSON(500, "Internal server error")
	}

	c.JSON(200, laundries)
}
