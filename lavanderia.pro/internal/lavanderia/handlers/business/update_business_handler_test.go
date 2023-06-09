package business

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
	createHandler := MakeCreateBusinessToUpdateHandler()
	updateHandler := MakeRegisterBusinessHandler()

	business, err := createHandler.Handle(&types.Business{
		Name: "test to update",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})

	businessUpdated, errUpdate := updateHandler.Handle(&types.Business{
		ID:   business.ID,
		Name: "test updated",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})

	assert.Nil(t, err, "Error on create business")
	assert.Nil(t, errUpdate, "Error on updated business")
	assert.NotEmpty(t, business, "Business is empty on create")
	assert.NotEmpty(t, businessUpdated, "Business is empty on delete")
	assert.NotEmpty(t, business.ID, "Business ID created is empty")
	assert.Equal(t, "test to update", business.Name, "Business name not created properly")
	assert.NotEmpty(t, business.Position, "Business position not created properly")
	assert.Equal(t, "test updated", businessUpdated.Name, "Business name not updated")
}

func MakeRegisterBusinessHandler() *UpdateBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	// repositoryAuth := repositories.NewAuthRepository(database)
	handler := NewUpdateBusinessHandler(repositoryBusiness)

	return handler
}

func MakeCreateBusinessToUpdateHandler() *CreateBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewBusinessRepository(database)
	handler := NewCreateBusinessHandler(repository)

	return handler
}
