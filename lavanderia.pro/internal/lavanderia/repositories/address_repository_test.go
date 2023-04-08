package repositories

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"strings"
	"testing"
	"time"
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

func TestUpdateAddress(t *testing.T) {
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
	assert.NotNil(t, address, "Create() returns nil result")
	assert.NotEmpty(t, address.CreatedAt, "CreatedAt is empty")

	address.Name = "UPDATED"

	addressUpdated, errUpdate := addressRepository.Update(&address)

	var addressUpdatedObject types.Address
	addressUpdatedObj, _ := bson.Marshal(addressUpdated)
	bson.Unmarshal(addressUpdatedObj, &addressUpdatedObject)

	assert.Equal(t, errUpdate, nil, "address Update() returns error")
	assert.NotEmpty(t, addressUpdated, "address Update() returns nil result")
	assert.Equal(t, "UPDATED", addressUpdatedObject.Name, "address name not save properly")
}

func TestGetAddresses(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	authRepository := MakeAuthRepositoryForTestGetAddress()
	clientRepository := MakeClientRepositoryForTestGetAddress()
	addressRepository := MakeAddressRepositoryToTest()

	// create random user
	pwd := []byte("PwD")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(pwd),
	}

	authCreated, errCreateAuth := authRepository.Create(auth)
	assert.Nil(t, errCreateAuth, "Error Auth returns not nil")
	assert.NotEmpty(t, authCreated, "auth created is empty")

	clientObj := &types.Client{
		Name: "Client Test",
		Auth: authCreated.ID,
	}

	clientRegistered, errClient := clientRepository.Create(clientObj)
	assert.Nil(t, errClient, "Error create client returns not nil")
	assert.NotEmpty(t, clientRegistered, "auth is empty")

	// Login Random User
	loginUser, errLogin := authRepository.GetByEmail(auth)
	assert.Nil(t, errLogin, "Error login client returns not nil")
	assert.NotEmpty(t, loginUser, "login user is empty")

	// Get Client By Auth
	client, errClient := clientRepository.GetClientByAuth(&types.Client{
		Auth: loginUser.ID,
	})
	assert.Equal(t, errClient, nil, "GetClientByAuth() returns error")
	assert.NotNil(t, client, "GetClientByAuth() returns nil result")

	// unmarshal client
	var clientObject types.Client
	clientUpdatedObj, _ := bson.Marshal(client)
	bson.Unmarshal(clientUpdatedObj, &clientObject)

	address1, errAdress1 := addressRepository.Create(&types.Address{
		Client: clientObject.ID,
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
		Name:    "TEST_ADDRESS_1",
		Extra:   "Call me",
		Phone:   "+123123123",
		Address: "Av. Pioneros 200, S.C Bariloche, Argentina",
	})

	assert.Equal(t, errAdress1, nil, "Create Address 1 returns error")
	assert.NotNil(t, address1, " Address 1 returns nil result")
	assert.NotEmpty(t, address1.CreatedAt, "address1 CreatedAt is empty")

	address2, errAddress2 := addressRepository.Create(&types.Address{
		Client: clientObject.ID,
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
		Name:    "TEST_ADDRESS_2",
		Extra:   "Call me",
		Phone:   "+123123123",
		Address: "Av. Pioneros 200, S.C Bariloche, Argentina",
	})

	assert.Equal(t, errAddress2, nil, "Create Address 2 returns error")
	assert.NotNil(t, address2, " Address 2 returns nil result")
	assert.NotEmpty(t, address2.CreatedAt, "address2 CreatedAt is empty")

	// Get all addresses len = 2
	adressesFound, errorFind := addressRepository.GetAddresses(&types.Address{
		Client: clientObject.ID,
	})

	assert.Equal(t, errorFind, nil, "GetAddresses() returns error")
	assert.NotNil(t, adressesFound, "GetAddresses() returns nil result")
	assert.Equal(t, 2, len(*adressesFound), "GetAddresses() different number of addresses created")
	fmt.Println(clientObject.ID)
}

func TestDeleteAddress(t *testing.T) {
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
	assert.NotNil(t, address, "Create() returns nil result")
	assert.NotEmpty(t, address.CreatedAt, "CreatedAt is empty")

	addressDeleted, errDelete := addressRepository.Delete(&address)

	assert.Equal(t, errDelete, nil, "address Delete() returns error")
	assert.NotEmpty(t, addressDeleted, "address Delete() returns nil result")
	assert.NotEmpty(t, addressDeleted.DeletedAt, "address DeletedAt returns null")
}

func MakeAddressRepositoryToTest() *AddressRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	addressRepository := NewAddressRepository(mongo)

	return addressRepository
}

func MakeAuthRepositoryForTestGetAddress() *AuthRepository {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := NewAuthRepository(database, config)

	return repositoryAuth
}

func MakeClientRepositoryForTestGetAddress() *ClientRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	clientRepository := NewClientRepository(mongo)

	return clientRepository
}
