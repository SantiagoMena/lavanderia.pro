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

func TestDeleteProductHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeCreateBusinessToDeleteHandler()
	deleteHandler := MakeDeleteProductHandler()

	product, err := createHandler.Handle(&types.Product{
		Name:  "test to delete",
		Price: 0.123,
	})

	productDeleted, errDel := deleteHandler.Handle(&types.Product{
		ID: product.ID,
	})

	assert.Nil(t, err, "Error on create business")
	assert.Nil(t, errDel, "Error on delete business")
	assert.NotEmpty(t, product, "Product is empty on create")
	assert.NotEmpty(t, productDeleted, "Product is empty on delete")
	assert.NotEmpty(t, productDeleted.ID, "Product ID created is empty")
	assert.NotEmpty(t, productDeleted.DeletedAt, "Product DeletedAt deleted is empty")
}

func MakeDeleteProductHandler() *DeleteProductHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewProductRepository(database)
	handler := NewDeleteProductHandler(repository)

	return handler
}

func MakeCreateBusinessToDeleteHandler() *CreateProductHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewProductRepository(database)
	handler := NewCreateProductHandler(repository)

	return handler
}
