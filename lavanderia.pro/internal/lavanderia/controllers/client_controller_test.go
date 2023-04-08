package controllers

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
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

func TestRegisterClient(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	// controllerAuth := MakeAuthControllerForTest()
	controllerClient := MakeClientControllerForTest()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	clientObj := &types.Client{
		Name: "test register",
	}

	client, err := controllerClient.RegisterClient(auth, clientObj)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, client, "Client is empty")
	assert.NotEmpty(t, client.ID, "Client ID is empty")
}

func TestGetClient(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controllerAuth := MakeAuthControllerForTest()
	controllerClient := MakeClientControllerForTest()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	clientObj := &types.Client{
		Name: "test register",
	}

	client, err := controllerClient.RegisterClient(auth, clientObj)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, client, "Client is empty")
	assert.NotEmpty(t, client.ID, "Client ID is empty")

	authClient, errAuth := controllerAuth.Login(auth)
	assert.Nil(t, errAuth, "Error returns not nil on Login")
	assert.NotEmpty(t, authClient, "AuthClient is empty on login")

	repositoryAuth := MakeAuthRepositoryForTestGetClient()
	authObject, errGetAuth := repositoryAuth.GetByEmail(auth)
	assert.Nil(t, errGetAuth, "Error returns not nil on getByEmail()")
	assert.NotEmpty(t, authObject, "authObject is empty on getByEmail()")

	clientGotten, errGet := controllerClient.GetClientByAuth(&types.Client{
		Auth: authObject.ID,
	})

	assert.Nil(t, errGet, "Error returns not nil on delete client")
	assert.NotEmpty(t, clientGotten, "Client is empty on delete client")
	assert.NotEmpty(t, clientGotten.ID, "Client ID is empty on delete client")
	assert.NotEmpty(t, clientGotten.CreatedAt, "Client CreatedAt is empty on delete client")
	assert.Equal(t, "test register", clientGotten.Name, "Name not get properly")
}

func TestPostClient(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controllerAuth := MakeAuthControllerForTest()
	controllerClient := MakeClientControllerForTest()
	repositoryAuth := MakeAuthRepositoryForTestGetClient()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	business := &types.Business{
		Name: "test register",
	}

	businessRegistered, err := controllerAuth.RegisterBusiness(auth, business)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, businessRegistered, "Business is empty")
	assert.NotEmpty(t, businessRegistered.ID, "Business ID is empty")

	authObject, errGetAuth := repositoryAuth.GetByEmail(auth)
	assert.Nil(t, errGetAuth, "Error returns not nil on getByEmail()")
	assert.NotEmpty(t, authObject, "authObject is empty on getByEmail()")

	clientPosted, errPost := controllerClient.PostClient(&types.Client{
		Auth: authObject.ID,
		Name: "POSTED",
	})

	assert.Nil(t, errPost, "Error returns not nil on post client")
	assert.NotEmpty(t, clientPosted, "Client is empty on post client")
	assert.NotEmpty(t, clientPosted.ID, "Client ID is empty on post client")
	assert.NotEmpty(t, clientPosted.CreatedAt, "Client CreatedAt is empty on post client")
	assert.Equal(t, "POSTED", clientPosted.Name, "Name not get properly")
}

func TestUpdateClient(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controllerAuth := MakeAuthControllerForTest()
	controllerClient := MakeClientControllerForTest()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	clientObj := &types.Client{
		Name: "test register",
	}

	client, err := controllerClient.RegisterClient(auth, clientObj)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, client, "Client is empty")
	assert.NotEmpty(t, client.ID, "Client ID is empty")

	authClient, errAuth := controllerAuth.Login(auth)
	assert.Nil(t, errAuth, "Error returns not nil on Login")
	assert.NotEmpty(t, authClient, "AuthClient is empty on login")

	repositoryAuth := MakeAuthRepositoryForTestGetClient()
	authObject, errGetAuth := repositoryAuth.GetByEmail(auth)
	assert.Nil(t, errGetAuth, "Error returns not nil on getByEmail()")
	assert.NotEmpty(t, authObject, "authObject is empty on getByEmail()")

	clientUpdated, errPut := controllerClient.UpdateClient(&types.Client{
		Auth: authObject.ID,
		Name: "UPDATED",
	})

	assert.Nil(t, errPut, "Error returns not nil on update client")
	assert.NotEmpty(t, clientUpdated, "Client is empty on update client")
	assert.NotEmpty(t, clientUpdated.ID, "Client ID is empty on update client")
	assert.NotEmpty(t, clientUpdated.CreatedAt, "Client CreatedAt is empty on update client")
	assert.Equal(t, "UPDATED", clientUpdated.Name, "Name not get properly")
}

func MakeClientControllerForTest() *ClientController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	authRepository := repositories.NewAuthRepository(database, config)
	clientRepository := repositories.NewClientRepository(database)

	controller := NewClientController(
		client.NewRegisterClientHandler(authRepository, clientRepository),
		client.NewGetClientHandler(authRepository, clientRepository),
		client.NewPostClientHandler(clientRepository),
		client.NewUpdateClientProfileHandler(clientRepository),
	)

	return controller
}

func MakeAuthControllerForTest() *AuthController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryClient := repositories.NewClientRepository(database)
	repositoryDelivery := repositories.NewDeliveryRepository(database)
	RegisterBusinessHandler := business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)
	RegisterClientHandler := client.NewRegisterClientHandler(repositoryAuth, repositoryClient)
	RegisterDeliveryHandler := delivery.NewRegisterDeliveryHandler(repositoryAuth, repositoryDelivery)
	LoginHandler := auth.NewLoginHandler(repositoryAuth, repositoryBusiness)
	RefreshTokenHandler := auth.NewRefreshTokenHandler(repositoryAuth)

	controller := NewAuthController(
		RegisterBusinessHandler,
		LoginHandler,
		RefreshTokenHandler,
		RegisterClientHandler,
		RegisterDeliveryHandler,
	)

	return controller
}

func MakeAuthRepositoryForTestGetClient() *repositories.AuthRepository {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)

	return repositoryAuth
}

func MakeBusinessControllerForTestClient() *BusinessController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)

	controller := NewBusinessController(
		business.NewGetAllBusinessHandler(repositoryBusiness),
		business.NewCreateBusinessHandler(repositoryBusiness),
		business.NewDeleteBusinessHandler(repositoryBusiness),
		business.NewUpdateBusinessHandler(repositoryBusiness),
		business.NewGetBusinessHandler(repositoryBusiness),
		business.NewGetAllBusinessByAuthHandler(repositoryBusiness),
	)

	return controller
}
