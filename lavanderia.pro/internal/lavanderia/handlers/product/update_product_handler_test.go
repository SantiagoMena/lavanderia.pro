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

func TestUpdateProductHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeCreateProductToUpdateHandler()
	updateHandler := MakeGetProductToTestUpdateHandler()

	product, err := createHandler.Handle(&types.Product{
		Name:  "test to update",
		Price: 0.123,
	})

	product.Price = 0.567
	product.Name = "UPDATED"

	productGotten, errGet := updateHandler.Handle(&product)

	assert.Nil(t, err, "Error on create product")
	assert.Nil(t, errGet, "Error on updated product")
	assert.NotEmpty(t, product, "Product is empty on create")
	assert.NotEmpty(t, productGotten, "Product is empty on get")
	assert.NotEmpty(t, product.ID, "Product ID update is empty")
	assert.Equal(t, "UPDATED", product.Name, "Product name not update properly")
	assert.Equal(t, 0.567, product.Price, "Product Price not update properly")
}

func MakeGetProductToTestUpdateHandler() *GetProductHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewProductRepository(database)
	handler := NewGetProductHandler(repository)

	return handler
}

func MakeCreateProductToUpdateHandler() *UpdateProductHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewProductRepository(database)
	handler := NewUpdateProductHandler(repository)

	return handler
}
