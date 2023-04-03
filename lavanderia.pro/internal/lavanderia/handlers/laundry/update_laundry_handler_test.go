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

func TestGetHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeCreateLaundryToUpdateHandler()
	updateHandler := MakeUpdateLaundryHandler()

	laundry, err := createHandler.Handle(&types.Laundry{
		Name: "test to update",
		Lat:  0.123,
		Long: 0.123,
	})

	laundryUpdated, errUpdate := updateHandler.Handle(&types.Laundry{
		Name: "test updated",
		Lat:  0.321,
		Long: 0.321,
	})

	assert.Nil(t, err, "Error on create laundry")
	assert.Nil(t, errUpdate, "Error on updated laundry")
	assert.NotEmpty(t, laundry, "Laundry is empty on create")
	assert.NotEmpty(t, laundryUpdated, "Laundry is empty on delete")
	assert.NotEmpty(t, laundry.ID, "Laundry ID created is empty")
	assert.Equal(t, "test to update", laundry.Name, "Laundry name not created properly")
	assert.Equal(t, 0.123, laundry.Lat, "Laundry lat not created properly")
	assert.Equal(t, 0.123, laundry.Long, "Laundry long not created properly")
	assert.Equal(t, "test updated", laundryUpdated.Name, "Laundry name not updated")
	assert.Equal(t, 0.321, laundryUpdated.Lat, "Laundry lat not updated")
	assert.Equal(t, 0.321, laundryUpdated.Long, "Laundry long not updated")
}

func MakeUpdateLaundryHandler() *UpdateLaundryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewLaundryRepository(database)
	handler := NewUpdateLaundryHandler(repository)

	return handler
}

func MakeCreateLaundryToUpdateHandler() *CreateLaundryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewLaundryRepository(database)
	handler := NewCreateLaundryHandler(repository)

	return handler
}
