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

func TestUpdateAddressHandler(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	createHandler := MakeCreateHandlerForTestUpdate()
	updateHandler := MakeUpdateHandlerForTestUpdate()

	address, err := createHandler.Handle(&types.Address{
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
		Name:    "TEST",
		Extra:   "Call me",
		Phone:   "+123123123",
		Address: "Av. Pioneros 201, S.C Bariloche, Argentina",
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, address, "FindAllBusiness() returns nil result")
	assert.NotEmpty(t, address.CreatedAt, "CreatedAt is empty")
	assert.Equal(t, "TEST", address.Name, "Name not save properly")

	address.Name = "UPDATED"
	addressUpdated, errUpdate := updateHandler.Handle(&address)
	assert.Equal(t, errUpdate, nil, "Create() returns error")
	assert.NotNil(t, addressUpdated, "update() returns nil result")
	assert.NotEmpty(t, addressUpdated.UpdatedAt, "UpdatedAt is empty")
	assert.Equal(t, "UPDATED", addressUpdated.Name, "Name not save properly")

}

func MakeCreateHandlerForTestUpdate() *CreateAddressHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewAddressRepository(database)
	handler := NewCreateAddressHandler(repository)

	return handler
}

func MakeUpdateHandlerForTestUpdate() *UpdateAddressHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewAddressRepository(database)
	handler := NewUpdateAddressHandler(repository)

	return handler
}
