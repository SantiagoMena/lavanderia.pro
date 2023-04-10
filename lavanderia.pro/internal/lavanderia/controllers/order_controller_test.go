package controllers

import (
	"fmt"
	"strings"
	"time"

	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/handlers/client"
	"lavanderia.pro/internal/lavanderia/handlers/delivery"
	"lavanderia.pro/internal/lavanderia/handlers/order"
	"lavanderia.pro/internal/lavanderia/repositories"
)

func TestPostOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "register order error")
	assert.NotEmpty(t, order, "register order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")
}

func TestGetOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderFound, errFind := orderController.GetOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errFind, "get order error")
	assert.NotEmpty(t, orderFound, "get order error")
	assert.NotEmpty(t, orderFound.CreatedAt, "CreatedAt order error")
}

func TestDeleteOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderDeleted, errDelete := orderController.DeleteOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errDelete, "delete order error")
	assert.NotEmpty(t, orderDeleted, "delete order error")
	assert.NotEmpty(t, orderDeleted.DeletedAt, "DeletedAt order error")
}

func TestAcceptOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderAccepted, errAccept := orderController.AcceptOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errAccept, "accept order error")
	assert.NotEmpty(t, orderAccepted, "accept order error")
	assert.NotEmpty(t, orderAccepted.AcceptedAt, "AcceptedAt order error")
}

func TestRejectOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderRejected, errReject := orderController.RejectOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errReject, "reject order error")
	assert.NotEmpty(t, orderRejected, "reject order error")
	assert.NotEmpty(t, orderRejected.RejectedAt, "RejectedAt order error")
}

func TestAssignPickUpOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderAccepted, errAccept := orderController.AcceptOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errAccept, "accept order error")
	assert.NotEmpty(t, orderAccepted, "accept order error")
	assert.NotEmpty(t, orderAccepted.AcceptedAt, "AcceptedAt order error")

	orderAssignPickUp, errAssignPickUp := orderController.AssignPickUpOrder(&types.Order{
		ID: order.ID,
		PickUp: types.Delivery{
			Name: "TEST PICKER UP",
		},
	})

	assert.Equal(t, nil, errAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp.AssignedPickUpAt, "assign pickup AssignedPickUpAt order error")
}

func TestPickUpOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderAccepted, errAccept := orderController.AcceptOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errAccept, "accept order error")
	assert.NotEmpty(t, orderAccepted, "accept order error")
	assert.NotEmpty(t, orderAccepted.AcceptedAt, "AcceptedAt order error")

	orderAssignPickUp, errAssignPickUp := orderController.AssignPickUpOrder(&types.Order{
		ID: order.ID,
		PickUp: types.Delivery{
			Name: "TEST PICKER UP",
		},
	})

	assert.Equal(t, nil, errAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp.AssignedPickUpAt, "assign pickup AssignedPickUpAt order error")

	orderPickUp, errPickUp := orderController.PickUpClientOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp.AssignedPickUpAt, "pickup AssignedPickUpAt order error")
}

func TestProcessOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderAccepted, errAccept := orderController.AcceptOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errAccept, "accept order error")
	assert.NotEmpty(t, orderAccepted, "accept order error")
	assert.NotEmpty(t, orderAccepted.AcceptedAt, "AcceptedAt order error")

	orderAssignPickUp, errAssignPickUp := orderController.AssignPickUpOrder(&types.Order{
		ID: order.ID,
		PickUp: types.Delivery{
			Name: "TEST PICKER UP",
		},
	})

	assert.Equal(t, nil, errAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp.AssignedPickUpAt, "assign pickup AssignedPickUpAt order error")

	orderPickUp, errPickUp := orderController.PickUpClientOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp.AssignedPickUpAt, "pickup AssignedPickUpAt order error")

	orderProcessing, errProcess := orderController.ProcessOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errProcess, "process order error")
	assert.NotEmpty(t, orderProcessing, "process order error")
	assert.NotEmpty(t, orderProcessing.ProcessingAt, "process ProcessingAt order error")
}

func TestFinishOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderAccepted, errAccept := orderController.AcceptOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errAccept, "accept order error")
	assert.NotEmpty(t, orderAccepted, "accept order error")
	assert.NotEmpty(t, orderAccepted.AcceptedAt, "AcceptedAt order error")

	orderAssignPickUp, errAssignPickUp := orderController.AssignPickUpOrder(&types.Order{
		ID: order.ID,
		PickUp: types.Delivery{
			Name: "TEST PICKER UP",
		},
	})

	assert.Equal(t, nil, errAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp.AssignedPickUpAt, "assign pickup AssignedPickUpAt order error")

	orderPickUp, errPickUp := orderController.PickUpClientOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp.AssignedPickUpAt, "pickup AssignedPickUpAt order error")

	orderProcessing, errProcess := orderController.ProcessOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errProcess, "process order error")
	assert.NotEmpty(t, orderProcessing, "process order error")
	assert.NotEmpty(t, orderProcessing.ProcessingAt, "process ProcessingAt order error")

	orderFinished, errFinish := orderController.FinishOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errFinish, "finish order error")
	assert.NotEmpty(t, orderFinished, "finish order error")
	assert.NotEmpty(t, orderFinished.FinishedAt, "finish FinishedAt order error")
}

func TestAssignDeliveryOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderAccepted, errAccept := orderController.AcceptOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errAccept, "accept order error")
	assert.NotEmpty(t, orderAccepted, "accept order error")
	assert.NotEmpty(t, orderAccepted.AcceptedAt, "AcceptedAt order error")

	orderAssignPickUp, errAssignPickUp := orderController.AssignPickUpOrder(&types.Order{
		ID: order.ID,
		PickUp: types.Delivery{
			Name: "TEST PICKER UP",
		},
	})

	assert.Equal(t, nil, errAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp.AssignedPickUpAt, "assign pickup AssignedPickUpAt order error")

	orderPickUp, errPickUp := orderController.PickUpClientOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp.AssignedPickUpAt, "pickup AssignedPickUpAt order error")

	orderProcessing, errProcess := orderController.ProcessOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errProcess, "process order error")
	assert.NotEmpty(t, orderProcessing, "process order error")
	assert.NotEmpty(t, orderProcessing.ProcessingAt, "process ProcessingAt order error")

	orderFinished, errFinish := orderController.FinishOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errFinish, "finish order error")
	assert.NotEmpty(t, orderFinished, "finish order error")
	assert.NotEmpty(t, orderFinished.FinishedAt, "finish FinishedAt order error")

	assignedDelivery, errAssignDelivery := orderController.AssignDeliveryOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errAssignDelivery, "assign delivery order error")
	assert.NotEmpty(t, assignedDelivery, "assign delivery order error")
	assert.NotEmpty(t, assignedDelivery.AssignedDeliveryAt, "assign delivery AssignedDeliveryAt order error")
}

func TestPickUpBusinessOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderAccepted, errAccept := orderController.AcceptOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errAccept, "accept order error")
	assert.NotEmpty(t, orderAccepted, "accept order error")
	assert.NotEmpty(t, orderAccepted.AcceptedAt, "AcceptedAt order error")

	orderAssignPickUp, errAssignPickUp := orderController.AssignPickUpOrder(&types.Order{
		ID: order.ID,
		PickUp: types.Delivery{
			Name: "TEST PICKER UP",
		},
	})

	assert.Equal(t, nil, errAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp.AssignedPickUpAt, "assign pickup AssignedPickUpAt order error")

	orderPickUp, errPickUp := orderController.PickUpClientOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp.AssignedPickUpAt, "pickup AssignedPickUpAt order error")

	orderProcessing, errProcess := orderController.ProcessOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errProcess, "process order error")
	assert.NotEmpty(t, orderProcessing, "process order error")
	assert.NotEmpty(t, orderProcessing.ProcessingAt, "process ProcessingAt order error")

	orderFinished, errFinish := orderController.FinishOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errFinish, "finish order error")
	assert.NotEmpty(t, orderFinished, "finish order error")
	assert.NotEmpty(t, orderFinished.FinishedAt, "finish FinishedAt order error")

	assignedDelivery, errAssignDelivery := orderController.AssignDeliveryOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errAssignDelivery, "assign delivery order error")
	assert.NotEmpty(t, assignedDelivery, "assign delivery order error")
	assert.NotEmpty(t, assignedDelivery.AssignedDeliveryAt, "assign delivery AssignedDeliveryAt order error")

	pickUpBusinessOrder, errPickUpBusiness := orderController.PickUpBusinessOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errPickUpBusiness, "pickup business order error")
	assert.NotEmpty(t, pickUpBusinessOrder, "pickup business order error")
	assert.NotEmpty(t, pickUpBusinessOrder.PickUpBusinessAt, "pickup business PickUpBusinessAt order error")
}

func TestDeliveryClientOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	orderController := MakeOrderControllerForTest()
	authController := MakeAuthForOrderController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authObject := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObject := &types.Business{
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, errBusiness := authController.RegisterBusiness(authObject, businessObject)

	assert.Equal(t, nil, errBusiness, "register business error")
	assert.NotEmpty(t, business, "register business error")

	order, errOrder := orderController.PostOrder(&types.Order{
		Business: business,
		Address:  types.Address{Name: "test"},
		Client:   types.Client{Name: "Client test"},
	})

	assert.Equal(t, nil, errOrder, "create order error")
	assert.NotEmpty(t, order, "create order error")
	assert.NotEmpty(t, order.CreatedAt, "CreatedAt order error")

	orderAccepted, errAccept := orderController.AcceptOrder(&types.Order{
		ID: order.ID,
	})

	assert.Equal(t, nil, errAccept, "accept order error")
	assert.NotEmpty(t, orderAccepted, "accept order error")
	assert.NotEmpty(t, orderAccepted.AcceptedAt, "AcceptedAt order error")

	orderAssignPickUp, errAssignPickUp := orderController.AssignPickUpOrder(&types.Order{
		ID: order.ID,
		PickUp: types.Delivery{
			Name: "TEST PICKER UP",
		},
	})

	assert.Equal(t, nil, errAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp, "assign pickup  order error")
	assert.NotEmpty(t, orderAssignPickUp.AssignedPickUpAt, "assign pickup AssignedPickUpAt order error")

	orderPickUp, errPickUp := orderController.PickUpClientOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp, "pickup order error")
	assert.NotEmpty(t, orderPickUp.AssignedPickUpAt, "pickup AssignedPickUpAt order error")

	orderProcessing, errProcess := orderController.ProcessOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errProcess, "process order error")
	assert.NotEmpty(t, orderProcessing, "process order error")
	assert.NotEmpty(t, orderProcessing.ProcessingAt, "process ProcessingAt order error")

	orderFinished, errFinish := orderController.FinishOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errFinish, "finish order error")
	assert.NotEmpty(t, orderFinished, "finish order error")
	assert.NotEmpty(t, orderFinished.FinishedAt, "finish FinishedAt order error")

	assignedDelivery, errAssignDelivery := orderController.AssignDeliveryOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errAssignDelivery, "assign delivery order error")
	assert.NotEmpty(t, assignedDelivery, "assign delivery order error")
	assert.NotEmpty(t, assignedDelivery.AssignedDeliveryAt, "assign delivery AssignedDeliveryAt order error")

	pickUpBusinessOrder, errPickUpBusiness := orderController.PickUpBusinessOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errPickUpBusiness, "pickup business order error")
	assert.NotEmpty(t, pickUpBusinessOrder, "pickup business order error")
	assert.NotEmpty(t, pickUpBusinessOrder.PickUpBusinessAt, "pickup business PickUpBusinessAt order error")

	deliveredClientOrder, errDeliveryClient := orderController.DeliveryClientOrder(&types.Order{ID: order.ID})

	assert.Equal(t, nil, errDeliveryClient, "delivery client order error")
	assert.NotEmpty(t, deliveredClientOrder, "delivery client order error")
	assert.NotEmpty(t, deliveredClientOrder.DeliveredClientAt, "delivery client DeliveredClientAt order error")
}

func MakeOrderControllerForTest() *OrderController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryOrder := repositories.NewOrderRepository(database)

	OrderController := NewOrderController(
		order.NewPostOrderHandler(repositoryOrder),
		order.NewGetOrderHandler(repositoryOrder),
		order.NewDeleteOrderHandler(repositoryOrder),
		order.NewAcceptOrderHandler(repositoryOrder),
		order.NewRejectOrderHandler(repositoryOrder),
		order.NewAssignPickUpOrderHandler(repositoryOrder),
		order.NewPickUpClientOrderHandler(repositoryOrder),
		order.NewProcessOrderHandler(repositoryOrder),
		order.NewFinishOrderHandler(repositoryOrder),
		order.NewAssignDeliveryOrderHandler(repositoryOrder),
		order.NewPickUpBusinessOrderHandler(repositoryOrder),
		order.NewDeliveryClientOrderHandler(repositoryOrder),
	)

	return OrderController
}

func MakeAuthForOrderController() *AuthController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryClient := repositories.NewClientRepository(database)
	repositoryDelivery := repositories.NewDeliveryRepository(database)
	RegisterBusinessHandler := business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)
	RegisterClientHandler := client.NewRegisterClientHandler(repositoryAuth, repositoryClient)
	RegisterDeliveryHandler := delivery.NewRegisterDeliveryHandler(repositoryAuth, repositoryDelivery)
	LoginHandler := auth.NewLoginHandler(repositoryAuth, repositoryBusiness)
	RefreshTokenHandler := auth.NewRefreshTokenHandler(repositoryAuth)

	controller := NewAuthController(
		RegisterBusinessHandler,
		LoginHandler,
		RefreshTokenHandler,
		RegisterClientHandler,
		RegisterDeliveryHandler,
	)

	return controller
}
