package controllers

import (
	"os"

	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"

	"github.com/gin-gonic/gin"
)

func Laundries(c *gin.Context) {
	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("MONGODB_DB")
	Mongodb := databases.NewMongoDatabase(uri, database)
	LaundryRepository := repositories.NewLaundryRepository(Mongodb)

	laundries, err := LaundryRepository.FindAllLaundries()
	if err != nil {
		c.JSON(500, "Internal server error")
	}

	c.JSON(200, laundries)
}
