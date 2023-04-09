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

func TestAssignPickUpOrderHandler(t *testing.T) {
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
	handler := MakePostOrderHandlerToTestAssignPickUp()

	order, errOrder := handler.Handle(&types.Order{
		Client:   *client,
		Address:  *address,
		Products: *productList,
		Business: *business,
	})

	assert.Equal(t, nil, errOrder, "error on create order handler")
	assert.NotEmpty(t, order, "order empty on create")
	assert.NotEmpty(t, order.CreatedAt, "order CreatedAt empty on get")

	// Accept order
	acceptOrderHandler := MakeAcceptOrderHandlerToTestAssignPickUp()
	acceptedOrder, errAccept := acceptOrderHandler.Handle(&order)

	assert.Equal(t, nil, errAccept, "error on accept order handler")
	assert.NotEmpty(t, acceptedOrder, "order empty on accept")
	assert.NotEmpty(t, acceptedOrder.AcceptedAt, "order AcceptedAt empty on get")

	// assignpickup order
	assignPickUpHandler := MakeAssignPickUpOrderHandlerToTestAssignPickUp()

	// Create Delivery
	orderAssignedPickUp, errAssignPickUp := assignPickUpHandler.Handle(&types.Order{
		ID: order.ID,
		PickUp: types.Delivery{
			Name: "pickup test",
		},
	})

	assert.Equal(t, nil, errAssignPickUp, "error on assign pickup order handler")
	assert.NotEmpty(t, orderAssignedPickUp, "order empty on assign pickup")
	assert.NotEmpty(t, orderAssignedPickUp.AssignedPickUpAt, "order AssignedPickUpAt empty on assign")
	// assert.NotEmpty(t, orderAssignedPickUp.PickUp, "order PickUp empty on assign")
}

func MakePostOrderHandlerToTestAssignPickUp() *PostOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewPostOrderHandler(repository)

	return handler
}

func MakeAssignPickUpOrderHandlerToTestAssignPickUp() *AssignPickUpOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewAssignPickUpOrderHandler(repository)

	return handler
}

func MakeAcceptOrderHandlerToTestAssignPickUp() *AcceptOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewAcceptOrderHandler(repository)

	return handler
}
