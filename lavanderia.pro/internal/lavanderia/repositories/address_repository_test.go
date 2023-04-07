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
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
		Name:    "House",
		Extra:   "Call me",
		Phone:   "+123123123",
		Address: "Av. Pioneros 200, S.C Bariloche, Argentina",
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, address, "FindAllBusiness() returns nil result")
	assert.NotEmpty(t, address.CreatedAt, "CreatedAt is empty")
}

func TestGetAddress(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	addressRepository := MakeAddressRepositoryToTest()

	address, err := addressRepository.Create(&types.Address{
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
		Name:    "TEST_ADDRESS",
		Extra:   "Call me",
		Phone:   "+123123123",
		Address: "Av. Pioneros 200, S.C Bariloche, Argentina",
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, address, "FindAllBusiness() returns nil result")
	assert.NotEmpty(t, address.CreatedAt, "CreatedAt is empty")

	addressFound, errFind := addressRepository.Get(&address)
	assert.Equal(t, errFind, nil, "address Get() returns error")
	assert.NotEmpty(t, addressFound, "address Get() returns nil result")
	assert.Equal(t, "TEST_ADDRESS", addressFound.Name, "address name not save properly")
}

func MakeAddressRepositoryToTest() *AddressRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	addressRepository := NewAddressRepository(mongo)

	return addressRepository
}
