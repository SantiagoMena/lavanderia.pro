package delivery

import (
	"github.com/joho/godotenv"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"

	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestPostDelivery(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	handler := MakePostDeliveryHandlerToTest()

	createHandler := MakeDeliveryRegisterHandlerToTestPostDelivery()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	deliveryObj := &types.Delivery{
		Name: "test register",
	}

	deliveryRegistered, errRegister := createHandler.Handle(auth, deliveryObj)

	assert.Nil(t, errRegister, "Error on register delivery")
	assert.NotEmpty(t, deliveryRegistered, "Delivery is empty on register")

	deliveryPosted, errPost := handler.Handle(&deliveryRegistered)
	assert.NotNil(t, errPost, "Not error handled on register twice delivery")
	assert.Empty(t, deliveryPosted, "Delivery registered twice")

	// TODO: Test create from auth without delivery
}

func MakePostDeliveryHandlerToTest() *PostDeliveryHandler {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	repository := repositories.NewDeliveryRepository(mongo)
	handler := NewPostDeliveryHandler(repository)

	return handler
}

func MakeDeliveryRegisterHandlerToTestPostDelivery() *RegisterDeliveryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryDelivery := repositories.NewDeliveryRepository(database)
	handler := NewRegisterDeliveryHandler(repositoryAuth, repositoryDelivery)

	return handler
}
