package controllers

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/address"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/handlers/client"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestCreateAddress(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controller := MakeAuthControllerForTestAddress()
	// authRepository := MakeAuthRepositoryForTestGetAddress()
	addressController := MakeAddressControllerForTest()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authLogin := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	clientObj := &types.Client{
		Name: "test register",
	}

	client, err := controller.RegisterClient(authLogin, clientObj)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, client, "Client registered is empty")

	addressCreated, errAddress := addressController.CreateAddress(&types.Address{
		Client:   client.ID,
		Position: []float64{-71.327767, -41.138444},
		Name:     "TEST",
		Extra:    "Call me",
		Address:  "Av. Pioneros 201, S.C Bariloche, Argentina",
	})

	assert.Nil(t, errAddress, "errAddress returns not nil")
	assert.NotEmpty(t, addressCreated, "addressCreated is empty")
	assert.Equal(t, "TEST", addressCreated.Name, "Address Name not save properly")
}

func MakeAuthControllerForTestAddress() *AuthController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryClient := repositories.NewClientRepository(database)
	controller := NewAuthController(
		business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness),
		auth.NewLoginHandler(repositoryAuth, repositoryBusiness),
		auth.NewRefreshTokenHandler(repositoryAuth),
		client.NewRegisterClientHandler(repositoryAuth, repositoryClient),
	)
	return controller
}

func MakeAddressControllerForTest() *AddressController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	addressRepository := repositories.NewAddressRepository(database)

	controller := NewAddressController(
		address.NewCreateAddressHandler(addressRepository),
	)

	return controller
}

func MakeAuthRepositoryForTestGetAddress() *repositories.AuthRepository {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)

	return repositoryAuth
}
