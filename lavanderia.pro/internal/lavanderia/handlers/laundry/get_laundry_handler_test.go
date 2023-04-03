package laundry

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"

	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
	"testing"
)

func TestUpdateHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeCreateLaundryToGetHandler()
	getHandler := MakeGetLaundryHandler()

	laundry, err := createHandler.Handle(&types.Laundry{
		Name: "test to update",
		Lat:  0.123,
		Long: 0.123,
	})

	laundryGotten, errGet := getHandler.Handle(&laundry)

	assert.Nil(t, err, "Error on create laundry")
	assert.Nil(t, errGet, "Error on updated laundry")
	assert.NotEmpty(t, laundry, "Laundry is empty on create")
	assert.NotEmpty(t, laundryGotten, "Laundry is empty on get")
	assert.NotEmpty(t, laundry.ID, "Laundry ID created is empty")
	assert.Equal(t, "test to update", laundry.Name, "Laundry name not created properly")
	assert.Equal(t, 0.123, laundry.Lat, "Laundry lat not created properly")
	assert.Equal(t, 0.123, laundry.Long, "Laundry long not created properly")
}

func MakeGetLaundryHandler() *GetLaundryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewLaundryRepository(database)
	handler := NewGetLaundryHandler(repository)

	return handler
}

func MakeCreateLaundryToGetHandler() *CreateLaundryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewLaundryRepository(database)
	handler := NewCreateLaundryHandler(repository)

	return handler
}
