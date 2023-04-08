package repositories

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"testing"
)

func TestCreateDelivery(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	delivery, err := NewDeliveryRepository(mongo).Create(&types.Delivery{
		Name: "test",
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, delivery, "FindAllDelivery() returns nil result")
	assert.NotEmpty(t, delivery.CreatedAt, "CreatedAt is empty")
}

func TestGetDeliveryByAuth(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	repository := NewDeliveryRepository(mongo)

	deliveryCreated, errCreate := repository.Create(&types.Delivery{
		Name: "TEST",
	})

	assert.Equal(t, errCreate, nil, "Error on create() delivery")
	assert.NotNil(t, deliveryCreated, "Create() returns nil result")
	assert.NotEmpty(t, deliveryCreated.CreatedAt, "CreatedAt is empty")
	assert.NotEmpty(t, deliveryCreated.ID, "ID created is empty")

	deliveryFound, errFind := repository.GetDeliveryByAuth(&deliveryCreated)

	assert.Equal(t, errFind, nil, "Error on GetDelivery() delivery")
	assert.NotNil(t, deliveryFound, "GetDelivery() returns nil result")
}
