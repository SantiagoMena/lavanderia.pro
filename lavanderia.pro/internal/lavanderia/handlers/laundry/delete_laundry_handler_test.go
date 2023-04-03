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

func TestDeleteHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeCreateLaundryToDeleteHandler()
	deleteHandler := MakeDeleteLaundryHandler()

	laundry, err := createHandler.Handle(&types.Laundry{
		Name: "test to delete",
		Lat:  0.123,
		Long: 0.123,
	})

	laundryDeleted, errDel := deleteHandler.Handle(&types.Laundry{
		ID: laundry.ID,
	})

	assert.Nil(t, err, "Error on create laundry")
	assert.Nil(t, errDel, "Error on delete laundry")
	assert.NotEmpty(t, laundry, "Laundry is empty on create")
	assert.NotEmpty(t, laundryDeleted, "Laundry is empty on delete")
	assert.NotEmpty(t, laundry.ID, "Laundry ID created is empty")
	assert.NotEmpty(t, laundry.ID, "Laundry ID deleted is empty")
}

func MakeDeleteLaundryHandler() *DeleteLaundryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewLaundryRepository(database)
	handler := NewDeleteLaundryHandler(repository)

	return handler
}

func MakeCreateLaundryToDeleteHandler() *CreateLaundryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewLaundryRepository(database)
	handler := NewCreateLaundryHandler(repository)

	return handler
}
