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

func TestUpdateHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeCreateBusinessToGetHandler()
	getHandler := MakeGetBusinessHandler()

	business, err := createHandler.Handle(&types.Business{
		Name: "test to update",
		Lat:  0.123,
		Long: 0.123,
	})

	businessGotten, errGet := getHandler.Handle(&business)

	assert.Nil(t, err, "Error on create business")
	assert.Nil(t, errGet, "Error on updated business")
	assert.NotEmpty(t, business, "Business is empty on create")
	assert.NotEmpty(t, businessGotten, "Business is empty on get")
	assert.NotEmpty(t, business.ID, "Business ID created is empty")
	assert.Equal(t, "test to update", business.Name, "Business name not created properly")
	assert.Equal(t, 0.123, business.Lat, "Business lat not created properly")
	assert.Equal(t, 0.123, business.Long, "Business long not created properly")
}

func MakeGetBusinessHandler() *GetBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewBusinessRepository(database)
	handler := NewGetBusinessHandler(repository)

	return handler
}

func MakeCreateBusinessToGetHandler() *CreateBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewBusinessRepository(database)
	handler := NewCreateBusinessHandler(repository)

	return handler
}
