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

func TestCreateClient(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	client, err := NewBusinessRepository(mongo).Create(&types.Business{
		Name: "test",
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, client, "Create() returns nil result")
	assert.NotEmpty(t, client.CreatedAt, "CreatedAt is empty")
}
