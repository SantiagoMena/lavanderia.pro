package delivery

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestGetDelivery(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeDeliveryRegisterHandlerToTestGet()
	getHandler := MakeGetDeliveryHandlerToTest()

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

	deliveryFound, errFind := getHandler.Handle(&deliveryRegistered)

	assert.Nil(t, errFind, "Error on getDelivery()")
	assert.NotEmpty(t, deliveryFound, "Delivery is empty on getDelivery()")

}

func MakeDeliveryRegisterHandlerToTestGet() *RegisterDeliveryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryDelivery := repositories.NewDeliveryRepository(database)
	handler := NewRegisterDeliveryHandler(repositoryAuth, repositoryDelivery)

	return handler
}

func MakeGetDeliveryHandlerToTest() *GetDeliveryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryDelivery := repositories.NewDeliveryRepository(database)
	handler := NewGetDeliveryHandler(repositoryDelivery)

	return handler
}
