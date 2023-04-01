package controllers

import (
	"lavanderia.pro/internal/lavanderia/repositories"

	"github.com/gin-gonic/gin"
)

func Laundries(c *gin.Context) {
	c.JSON(200, repositories.FindAllLaundries())
}
