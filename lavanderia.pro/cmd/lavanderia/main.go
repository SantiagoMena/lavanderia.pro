package main

import (
	"log"

	"github.com/joho/godotenv"

	"lavanderia.pro/internal/lavanderia/routers"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}

	r := routers.SetupRouter()

	r.Run()
}
