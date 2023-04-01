package main

import (
	"github.com/gin-gonic/gin"

	"lavanderia.pro/internal/lavanderia/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/laundries", controllers.Laundries)

	r.Run()
}
