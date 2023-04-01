package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"lavanderia.pro/internal/lavanderia/controllers"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}

	r := gin.Default()
	r.GET("/laundries", controllers.Laundries)
	r.GET("/ping", Ping)

	r.Run()
}

func Ping(c *gin.Context) {
	type status struct {
		Status string `json:"status"`
	}

	statusObj := status{
		Status: "ok",
	}

	c.JSON(200, statusObj)
}
