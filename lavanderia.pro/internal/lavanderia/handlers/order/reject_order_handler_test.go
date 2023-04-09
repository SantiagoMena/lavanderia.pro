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

func TestRejectOrderHandler(t *testing.T) {
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
	handler := MakePostOrderHandlerToTestReject()
	order, errOrder := handler.Handle(&types.Order{
		Client:   *client,
		Address:  *address,
		Products: *productList,
		Business: *business,
	})

	assert.Equal(t, nil, errOrder, "error on create order handler")
	assert.NotEmpty(t, order, "order empty on create")
	assert.NotEmpty(t, order.CreatedAt, "order CreatedAt empty on get")

	// reject order
	rejectHandler := MakeRejectOrderHandlerToTestReject()
	orderRejected, errReject := rejectHandler.Handle(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errReject, "error on reject order handler")
	assert.NotEmpty(t, orderRejected, "order empty on reject")
	assert.NotEmpty(t, orderRejected.RejectedAt, "order RejectedAt empty on get")
}

func MakePostOrderHandlerToTestReject() *PostOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewPostOrderHandler(repository)

	return handler
}

func MakeRejectOrderHandlerToTestReject() *RejectOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewRejectOrderHandler(repository)

	return handler
}
