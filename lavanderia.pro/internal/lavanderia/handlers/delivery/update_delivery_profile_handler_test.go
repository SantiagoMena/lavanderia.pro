package delivery

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
	"lavanderia.pro/internal/lavanderia/repositories"
)

func TestUpdateDeliveryProfileHandler(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	updateHandler := MakeUpdateHandlerToTest()
	registerHandler := MakeDeliveryRegisterHandlerToTestUpdateDelivery()

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

	deliveryRegistered, errRegister := registerHandler.Handle(auth, deliveryObj)

	assert.Nil(t, errRegister, "Error on register delivery")
	assert.NotEmpty(t, deliveryRegistered, "Delivery is empty on register")

	deliveryRegistered.Name = "UPDATED"
	deliveryUpdated, errUpdate := updateHandler.Handle(&deliveryRegistered)

	assert.Nil(t, errUpdate, "Error on update delivery")
	assert.NotEmpty(t, deliveryRegistered, "Delivery is empty on update")
	assert.Equal(t, "UPDATED", deliveryUpdated.Name, "Delivery name updated not properly")

}

func MakeUpdateHandlerToTest() *UpdateDeliveryProfileHandler {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	repository := repositories.NewDeliveryRepository(mongo)
	handler := NewUpdateDeliveryProfileHandler(repository)

	return handler
}

func MakeDeliveryRegisterHandlerToTestUpdateDelivery() *RegisterDeliveryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryDelivery := repositories.NewDeliveryRepository(database)
	handler := NewRegisterDeliveryHandler(repositoryAuth, repositoryDelivery)

	return handler
}
