package controllers

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	type status struct {
		Status string `json:"status"`
	}

	statusObj := status{
		Status: "ok",
	}

	c.JSON(200, statusObj)
}
