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

func TestGetOrderHandler(t *testing.T) {
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

	// create product
	product := &types.Product{
		Name:  "product test",
		Price: 0.123,
	}

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
		{Product: *product, Amount: 10},
	}

	// create order
	handler := MakePostOrderHandlerToTestGet()
	order, errOrder := handler.Handle(&types.Order{
		Client:   *client,
		Address:  *address,
		Products: *productList,
		Business: *business,
	})

	assert.Equal(t, nil, errOrder, "error on create order handler")
	assert.NotEmpty(t, order, "order empty on create")
	assert.NotEmpty(t, order.CreatedAt, "order CreatedAt empty on get")
}

func MakePostOrderHandlerToTestGet() *PostOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	productRepository := repositories.NewProductRepository(database)
	handler := NewPostOrderHandler(repository, productRepository)

	return handler
}
