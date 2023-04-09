package business

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"errors"
	"strings"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
)

func TestRegisterBusinessDeliveryHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	registerDeliveryHandler := MakeCreateBusinessDeliveryToRegisterHandler()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObj := &types.Business{
		Name: "test register",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	deliveryObj := &types.Delivery{
		Name: "test register",
	}

	deliveryRegistered, errRegister := registerDeliveryHandler.Handle(auth, businessObj, deliveryObj)

	assert.Nil(t, errRegister, "Error on register delivery")
	assert.NotEmpty(t, deliveryRegistered, "delivery is empty on register")

	deliveryRegisteredSameEmail, errRegisterSame := registerDeliveryHandler.Handle(auth, businessObj, deliveryObj)
	assert.Equal(t, errRegisterSame, errors.New("auth already exists"), "Error auth already exists not work")
	assert.Empty(t, deliveryRegisteredSameEmail, "Error auth already exists return empty delivery")

}

func MakeCreateBusinessDeliveryToRegisterHandler() *RegisterBusinessDeliveryHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryDelivery := repositories.NewDeliveryRepository(database)
	handler := NewRegisterBusinessDeliveryHandler(repositoryAuth, repositoryBusiness, repositoryDelivery)

	return handler
}
