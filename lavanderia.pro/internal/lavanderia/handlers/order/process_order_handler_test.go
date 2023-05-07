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

func TestProcessClientOrderHandler(t *testing.T) {
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
	handler := MakePostOrderHandlerToTestProcess()

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

	acceptOrderHandler := MakeAcceptOrderHandlerToTestProcess()
	acceptedOrder, errAcceptOrder := acceptOrderHandler.Handle(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errAcceptOrder, "error on accept order handler")
	assert.NotEmpty(t, acceptedOrder, "order empty on accept")
	assert.NotEmpty(t, acceptedOrder.AcceptedAt, "order AcceptedAt empty on accept")

	// assignpickup
	assignPickUpOrderHandler := MakeAssignPickUpOrderHandlerToTestAssignPickUp()

	// Create Delivery
	pickup := &types.Delivery{
		Name: "pickup test",
	}

	// Pickup Client
	orderAssignedPickUp, errPickUpClient := assignPickUpOrderHandler.Handle(&types.Order{
		ID:     acceptedOrder.ID,
		PickUp: *pickup,
	})

	fmt.Println(acceptedOrder.ID)
	assert.Equal(t, nil, errPickUpClient, "error on assign pickup order handler")
	assert.NotEmpty(t, orderAssignedPickUp, "order empty on assign pickup")
	assert.NotEmpty(t, orderAssignedPickUp.AssignedPickUpAt, "order AssignedPickUpAt empty on assign")
	// TODO: Check order assign pickup
	assert.NotEmpty(t, orderAssignedPickUp.PickUp, "order PickUp empty on assign")

	pickUpClientHandler := MakePickUpClientOrderHandlerToTestProcess()

	orderPicketUp, errPickUpClient := pickUpClientHandler.Handle(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errPickUpClient, "error on pickup order handler")
	assert.NotEmpty(t, orderPicketUp, "order empty on pickup")
	assert.NotEmpty(t, orderPicketUp.PickUpClientAt, "order PickUpClientAt empty on pickup")

	processHandler := MakeProcessOrderHandlerToTestProcess()

	orderProcess, errProcessClient := processHandler.Handle(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errProcessClient, "error on process order handler")
	assert.NotEmpty(t, orderProcess, "order empty on process")
	assert.NotEmpty(t, orderProcess.ProcessingAt, "order ProcessingAt empty on process")
}

func MakePostOrderHandlerToTestProcess() *PostOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	productRepository := repositories.NewProductRepository(database)
	handler := NewPostOrderHandler(repository, productRepository)

	return handler
}

func MakeAssignPickUpOrderHandlerToTestProcess() *AssignPickUpOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewAssignPickUpOrderHandler(repository)

	return handler
}

func MakePickUpOrderHandlerToTestProcess() *PickUpClientOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewPickUpClientOrderHandler(repository)

	return handler
}

func MakePickUpClientOrderHandlerToTestProcess() *PickUpClientOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewPickUpClientOrderHandler(repository)

	return handler
}

func MakeAcceptOrderHandlerToTestProcess() *AcceptOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewAcceptOrderHandler(repository)

	return handler
}

func MakeProcessOrderHandlerToTestProcess() *ProcessOrderHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewOrderRepository(database)
	handler := NewProcessOrderHandler(repository)

	return handler
}
