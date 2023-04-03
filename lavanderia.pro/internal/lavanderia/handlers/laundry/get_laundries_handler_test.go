package laundry

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"lavanderia.pro/internal/lavanderia/config"

	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
	"testing"
)

func TestGetLaundriesHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	handler := MakeGetLaundriesHandler()

	laundries, err := handler.Handle()

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, laundries, "Laundry is empty")
}

func MakeGetLaundriesHandler() *GetLaundriesHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewLaundryRepository(database)
	handler := NewGetLaundriesHandler(repository)

	return handler
}
