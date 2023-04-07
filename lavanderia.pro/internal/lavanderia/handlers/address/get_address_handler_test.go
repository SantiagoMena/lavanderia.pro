package address

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

func TestGetAddress(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeCreateAddressToGetHandler()
	getHandler := MakeGetAddressToTestGetHandler()

	address, err := createHandler.Handle(&types.Address{
		Name: "test_to_get",
	})

	addressGotten, errGet := getHandler.Handle(&address)

	assert.Nil(t, err, "Error on create address")
	assert.Nil(t, errGet, "Error on updated address")
	assert.NotEmpty(t, address, "Address is empty on create")
	assert.NotEmpty(t, addressGotten, "Address is empty on get")
	assert.NotEmpty(t, address.ID, "Address ID created is empty")
	assert.Equal(t, "test_to_get", address.Name, "Address Name not created properly")

}

func MakeGetAddressToTestGetHandler() *GetAddressHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewAddressRepository(database)
	handler := NewGetAddressHandler(repository)

	return handler
}

func MakeCreateAddressToGetHandler() *CreateAddressHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewAddressRepository(database)
	handler := NewCreateAddressHandler(repository)

	return handler
}
