package order

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
)

func TestDeleteOrderHandler(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	// create business
	business := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	productRepository := MakeProductRepositoryToTestDeleteOrder()
	// create product
	productObject := &types.Product{
		Name:  "product test",
		Price: 0.123,
	}
	product, errCreateProduct := productRepository.Create(productObject)

	assert.Equal(t, nil, errCreateProduct, "error on create product handler")
	assert.NotEmpty(t, product, "product empty on create")

	// create client
	client := &types.Client{
		Name: "client test",
	}

	// create address
	address := &types.Address{
		Name: "Home Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	// create product list
	productList := &[]types.OrderProduct{
		{Product: types.Product{ID: product.ID}, Amount: 10},
	}

	// create order
	handler := MakePostOrderHandlerToTestDelete()
	order, errOrder := handler.Handle(&types.Order{
		Client:   *client,
		Address:  *address,
		Products: *productList,
		Business: *business,
	})

	assert.Equal(t, nil, errOrder, "error on create order handler")
	assert.NotEmpty(t, order, "order empty on create")
	assert.NotEmpty(t, order.CreatedAt, "order CreatedAt empty on get")

	// delete order
	deleteHandler := MakeDeleteOrderHandlerToTestDelete()
	orderDeleted, errDelete := deleteHandler.Handle(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errDelete, "error on delete order handler")
	assert.NotEmpty(t, orderDeleted, "order empty on delete")
	assert.NotEmpty(t, orderDeleted.DeletedAt, "order DeletedAt empty on get")
}

func MakePostOrderHandlerToTestDelete() *PostOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	productRepository := repositories.NewProductRepository(database)
	handler := NewPostOrderHandler(repository, productRepository)

	return handler
}

func MakeDeleteOrderHandlerToTestDelete() *DeleteOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewDeleteOrderHandler(repository)

	return handler
}

func MakeProductRepositoryToTestDeleteOrder() *repositories.ProductRepository {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewProductRepository(database)

	return repository
}
