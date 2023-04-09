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

	orderController := MakeAOrderControllerForTest()
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

func MakeAOrderControllerForTest() *OrderController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryOrder := repositories.NewOrderRepository(database)
	postOrderHandler := order.NewPostOrderHandler(repositoryOrder)
	OrderController := NewOrderController(postOrderHandler)

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
