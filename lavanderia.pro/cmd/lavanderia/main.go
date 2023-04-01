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

	r.Run()
}
