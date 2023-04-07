package repositories

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	addressRepository := MakeAddressRepositoryToTest()

	address, err := addressRepository.Create(&types.Address{
		Position: []float64{-71.327767, -41.138444},
		Name:     "House",
		Extra:    "Call me",
		Phone:    "+123123123",
		Address:  "Av. Pioneros 200, S.C Bariloche, Argentina",
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, address, "FindAllBusiness() returns nil result")
	assert.NotEmpty(t, address.CreatedAt, "CreatedAt is empty")
}

func MakeAddressRepositoryToTest() *AddressRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	addressRepository := NewAddressRepository(mongo)

	return addressRepository
}
