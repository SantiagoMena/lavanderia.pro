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
