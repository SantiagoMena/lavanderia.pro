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

func TestCreateAddressHandler(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	handler := MakeCreateHandlerForTest()

	address, err := handler.Handle(&types.Address{
		Position: []float64{-71.327767, -41.138444},
		Name:     "TEST",
		Extra:    "Call me",
		Phone:    "+123123123",
		Address:  "Av. Pioneros 201, S.C Bariloche, Argentina",
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, address, "FindAllBusiness() returns nil result")
	assert.NotEmpty(t, address.CreatedAt, "CreatedAt is empty")
	assert.Equal(t, "TEST", address.Name, "Name not save properly")
}

func MakeCreateHandlerForTest() *CreateAddressHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewAddressRepository(database)
	handler := NewCreateAddressHandler(repository)

	return handler
}
