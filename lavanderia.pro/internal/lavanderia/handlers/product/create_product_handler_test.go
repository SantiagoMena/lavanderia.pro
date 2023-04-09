package product

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"

	"testing"

	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
)

func TestCreateProductHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	handler := MakeCreateProductHandler()

	product, err := handler.Handle(&types.Product{
		Name:  "test",
		Price: 0.123,
	})

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, product, "Product is empty")
	assert.NotEmpty(t, product.ID, "Product ID is empty")
	assert.NotEmpty(t, product.CreatedAt, "Product CreatedAt is empty")
}

func MakeCreateProductHandler() *CreateProductHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewProductRepository(database)
	handler := NewCreateProductHandler(repository)

	return handler
}
