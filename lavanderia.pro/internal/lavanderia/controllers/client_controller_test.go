package controllers

import (
	"fmt"

	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/handlers/client"
	"lavanderia.pro/internal/lavanderia/repositories"
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

func MakeClientControllerForTest() *ClientController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	authRepository := repositories.NewAuthRepository(database, config)
	clientRepository := repositories.NewClientRepository(database)

	controller := NewClientController(
		client.NewRegisterClientHandler(authRepository, clientRepository),
		client.NewGetClientHandler(authRepository, clientRepository),
		client.NewPostClientHandler(clientRepository),
	)

	return controller
}

func MakeAuthControllerForTest() *AuthController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryClient := repositories.NewClientRepository(database)
	RegisterBusinessHandler := business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)
	RegisterClientHandler := client.NewRegisterClientHandler(repositoryAuth, repositoryClient)
	LoginHandler := auth.NewLoginHandler(repositoryAuth, repositoryBusiness)
	RefreshTokenHandler := auth.NewRefreshTokenHandler(repositoryAuth)

	controller := NewAuthController(
		RegisterBusinessHandler,
		LoginHandler,
		RefreshTokenHandler,
		RegisterClientHandler,
	)

	return controller
}

func MakeAuthRepositoryForTestGetClient() *repositories.AuthRepository {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)

	return repositoryAuth
}
