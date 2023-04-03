package business

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"lavanderia.pro/internal/lavanderia/config"

	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
	"testing"
)

func TestGetAllBusinessHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	handler := MakeGetAllBusinessHandler()

	allBusiness, err := handler.Handle()

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, allBusiness, "Business is empty")
}

func MakeGetAllBusinessHandler() *GetAllBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewBusinessRepository(database)
	handler := NewGetAllBusinessHandler(repository)

	return handler
}
