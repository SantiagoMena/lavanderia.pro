package address

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/controllers"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/handlers/client"
	"lavanderia.pro/internal/lavanderia/handlers/delivery"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestGetAddresses(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	// Register Client
	authController := MakeAuthControllerToTestGetAddresses()

	pwd := []byte("PwD")

	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(pwd),
	}

	clientObj := &types.Client{
		Name: "Client Test",
	}

	client, errClient := authController.RegisterClient(auth, clientObj)
	assert.Nil(t, errClient, "error on register client")
	assert.NotNil(t, client, "register client return nil")

	clientRepository := MakeClientRepositoryForTestGetAddresses()

	// Get Client Id from auth
	clientAuth, errClientAuth := clientRepository.GetClientByAuth(&client)
	assert.Equal(t, errClientAuth, nil, "GetClientByAuth() returns error")
	assert.NotNil(t, clientAuth, "GetClientByAuth() returns nil result")

	// unmarshal client
	var clientObject types.Client
	clientUpdatedObj, _ := bson.Marshal(client)
	bson.Unmarshal(clientUpdatedObj, &clientObject)

	// Create Address 1
	addressRepository := MakeAddressRepositoryToTestGetAddresses()

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

	// Find Addresses

	// Check 2 Addresses

}

func MakeAuthControllerToTestGetAddresses() *controllers.AuthController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	authRepository := repositories.NewAuthRepository(database, config)
	clientRepository := repositories.NewClientRepository(database)
	deliveryRepository := repositories.NewDeliveryRepository(database)
	businessRepository := repositories.NewBusinessRepository(database)
	RegisterBusinessHandler := business.NewRegisterBusinessHandler(authRepository, businessRepository)
	LoginHandler := auth.NewLoginHandler(authRepository, businessRepository)
	RefreshTokenHandler := auth.NewRefreshTokenHandler(authRepository)
	RegisterClientHandler := client.NewRegisterClientHandler(authRepository, clientRepository)
	RegisterDeliveryHandler := delivery.NewRegisterDeliveryHandler(authRepository, deliveryRepository)

	AuthController := controllers.NewAuthController(
		RegisterBusinessHandler,
		LoginHandler,
		RefreshTokenHandler,
		RegisterClientHandler,
		RegisterDeliveryHandler,
	)

	return AuthController
}

func MakeAddressRepositoryToTestGetAddresses() *repositories.AddressRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	addressRepository := repositories.NewAddressRepository(mongo)

	return addressRepository
}

func MakeClientRepositoryForTestGetAddresses() *repositories.ClientRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	clientRepository := repositories.NewClientRepository(mongo)

	return clientRepository
}
