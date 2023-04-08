package delivery

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"errors"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestRegisterDeliveryHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeCreateDeliveryToRegisterHandler()
	// updateHandler := MakeUpdateDeliveryHandler()

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

	deliveryRegisteredSameEmail, errRegisterSame := createHandler.Handle(auth, deliveryObj)
	assert.Equal(t, errRegisterSame, errors.New("auth already exists"), "Error auth already exists not work")
	assert.Equal(t, types.Delivery{}, deliveryRegisteredSameEmail, "Error auth already exists return empty delivery")

}

// func MakeUpdateDeliveryHandler() *UpdateDeliveryHandler {
// 	config := config.NewConfig()
// 	database := databases.NewMongoDatabase(config)
// 	repository := repositories.NewDeliveryRepository(database)
// 	handler := NewUpdateDeliveryHandler(repository)

// 	return handler
// }

func MakeCreateDeliveryToRegisterHandler() *RegisterDeliveryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryDelivery := repositories.NewDeliveryRepository(database)
	handler := NewRegisterDeliveryHandler(repositoryAuth, repositoryDelivery)

	return handler
}
