package controllers

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/laundry"
	"lavanderia.pro/internal/lavanderia/repositories"
	"testing"
)

func TestGetLaundries(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controller := MakeController()
	laundries, err := controller.GetLaundries()

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, laundries, "Laundries is empty")
}

func TestPostLaundries(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controller := MakeController()

	laundry, err := controller.PostLaundry(&types.Laundry{
		Name: "test",
		Lat:  0.123,
		Long: 0.123,
	})

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, laundry, "Laundry is empty")
	assert.NotEmpty(t, laundry.ID, "Laundry ID is empty")
}

func MakeController() *LaundryController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewLaundryRepository(database)
	controller := NewLaundryController(laundry.NewGetLaundriesHandler(repository), laundry.NewCreateLaundryHandler(repository))
	return controller
}
