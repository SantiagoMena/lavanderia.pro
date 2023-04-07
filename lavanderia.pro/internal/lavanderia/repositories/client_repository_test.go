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

	client, err := NewClientRepository(mongo).Create(&types.Client{
		Name: "test",
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, client, "Create() returns nil result")
	assert.NotEmpty(t, client.CreatedAt, "CreatedAt is empty")
}

func TestGetClient(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	repository := NewClientRepository(mongo)

	clientCreated, errCreate := repository.Create(&types.Client{
		Name: "TEST",
	})

	assert.Equal(t, errCreate, nil, "Error on create() client")
	assert.NotNil(t, clientCreated, "Create() returns nil result")
	assert.NotEmpty(t, clientCreated.CreatedAt, "CreatedAt is empty")
	assert.NotEmpty(t, clientCreated.ID, "ID created is empty")

	fmt.Println("clientCreated")
	fmt.Println(clientCreated)

	clientFound, errFind := repository.GetClientByAuth(&clientCreated)
	fmt.Println("clientFound")
	fmt.Println(clientFound)

	assert.Equal(t, errFind, nil, "Error on GetClient() client")
	assert.NotNil(t, clientFound, "GetClient() returns nil result")
	// assert.Equal(t, "TEST", clientFound.Name, "Error on GetClient() client")
}
