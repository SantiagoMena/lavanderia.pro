package product

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"

	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
	"testing"
)

func TestUpdateHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeCreateProductToGetHandler()
	getHandler := MakeGetProductToTestGetHandler()

	product, err := createHandler.Handle(&types.Product{
		Name:  "test to update",
		Price: 0.123,
	})

	productGotten, errGet := getHandler.Handle(&product)

	assert.Nil(t, err, "Error on create product")
	assert.Nil(t, errGet, "Error on updated product")
	assert.NotEmpty(t, product, "Product is empty on create")
	assert.NotEmpty(t, productGotten, "Product is empty on get")
	assert.NotEmpty(t, product.ID, "Product ID created is empty")
	assert.Equal(t, "test to update", product.Name, "Product name not created properly")
	assert.Equal(t, 0.123, product.Price, "Product Price not created properly")
}

func MakeGetProductToTestGetHandler() *GetProductHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewProductRepository(database)
	handler := NewGetProductHandler(repository)

	return handler
}

func MakeCreateProductToGetHandler() *CreateProductHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewProductRepository(database)
	handler := NewCreateProductHandler(repository)

	return handler
}
