package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"lavanderia.pro/internal/lavanderia/controllers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Ping)
	r.GET("/laundries", controllers.Laundries)

	return r
}

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}

	r := setupRouter()

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
