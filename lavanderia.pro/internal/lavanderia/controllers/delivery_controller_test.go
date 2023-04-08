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

func TestPostDelivery(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controllerAuth := MakeAuthControllerForTestDelivery()
	controllerDelivery := MakeDeliveryControllerForTestDelivery()
	repositoryAuth := MakeAuthRepositoryForTestGetDelivery()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	client := &types.Client{
		Name: "test register",
	}

	clientRegistered, errClient := controllerAuth.RegisterClient(auth, client)

	assert.Nil(t, errClient, "Error returns not nil")
	assert.NotEmpty(t, clientRegistered, "Client is empty")
	assert.NotEmpty(t, clientRegistered.ID, "Client ID is empty")

	authObject, errGetAuth := repositoryAuth.GetByEmail(auth)
	assert.Nil(t, errGetAuth, "Error returns not nil on getByEmail()")
	assert.NotEmpty(t, authObject, "authObject is empty on getByEmail()")

	deliveryPosted, errPost := controllerDelivery.PostDelivery(&types.Delivery{
		Auth: authObject.ID,
		Name: "POSTED",
	})

	assert.Nil(t, errPost, "Error returns not nil on post delivery")
	assert.NotEmpty(t, deliveryPosted, "Delivery is empty on post delivery")
	assert.NotEmpty(t, deliveryPosted.ID, "Delivery ID is empty on post delivery")
	assert.NotEmpty(t, deliveryPosted.CreatedAt, "Delivery CreatedAt is empty on post delivery")
	assert.Equal(t, "POSTED", deliveryPosted.Name, "Name not get properly")
}

func MakeAuthRepositoryForTestGetDelivery() *repositories.AuthRepository {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)

	return repositoryAuth
}

func MakeDeliveryControllerForTestDelivery() *DeliveryController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryDelivery := repositories.NewDeliveryRepository(database)

	controller := NewDeliveryController(
		delivery.NewPostDeliveryHandler(repositoryDelivery),
		delivery.NewGetDeliveryHandler(repositoryDelivery),
	)

	return controller
}

func MakeAuthControllerForTestDelivery() *AuthController {
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
