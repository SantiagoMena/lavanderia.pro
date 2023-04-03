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

func TestDeleteHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeCreateBusinessToDeleteHandler()
	deleteHandler := MakeDeleteBusinessHandler()

	business, err := createHandler.Handle(&types.Business{
		Name: "test to delete",
		Lat:  0.123,
		Long: 0.123,
	})

	businessDeleted, errDel := deleteHandler.Handle(&types.Business{
		ID: business.ID,
	})

	assert.Nil(t, err, "Error on create business")
	assert.Nil(t, errDel, "Error on delete business")
	assert.NotEmpty(t, business, "Business is empty on create")
	assert.NotEmpty(t, businessDeleted, "Business is empty on delete")
	assert.NotEmpty(t, business.ID, "Business ID created is empty")
	assert.NotEmpty(t, business.ID, "Business ID deleted is empty")
}

func MakeDeleteBusinessHandler() *DeleteBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewBusinessRepository(database)
	handler := NewDeleteBusinessHandler(repository)

	return handler
}

func MakeCreateBusinessToDeleteHandler() *CreateBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewBusinessRepository(database)
	handler := NewCreateBusinessHandler(repository)

	return handler
}
