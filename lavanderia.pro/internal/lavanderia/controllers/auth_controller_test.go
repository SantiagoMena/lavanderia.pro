package controllers

import (
	"fmt"

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
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestRegisterBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controller := MakeAuthController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authRegister := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObj := &types.Business{
		Name: "test register",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, err := controller.RegisterBusiness(authRegister, businessObj)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, business, "Business is empty")
}

func TestLogin(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controller := MakeAuthController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authLogin := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObj := &types.Business{
		Name: "test register",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, err := controller.RegisterBusiness(authLogin, businessObj)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, business, "Business is empty")

	businessLogged, errLogin := controller.Login(authLogin)

	assert.Nil(t, errLogin, "Error Login returns not nil")
	assert.NotEmpty(t, businessLogged, "Business Logged is empty")
}

func TestRefreshToken(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controller := MakeAuthController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authLogin := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObj := &types.Business{
		Name: "test register",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	business, err := controller.RegisterBusiness(authLogin, businessObj)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, business, "Business is empty")

	businessLogged, errLogin := controller.Login(authLogin)

	assert.Nil(t, errLogin, "Error Login returns not nil")
	assert.NotEmpty(t, businessLogged, "Business Logged is empty")
	assert.NotEmpty(t, businessLogged.Token, "Business Logged Token is empty")

	RefreshToken, errRefresh := controller.RefreshToken(businessLogged.RefreshToken)

	assert.Nil(t, errRefresh, "Error Login returns not nil")
	assert.NotEmpty(t, RefreshToken, "RefreshToken is empty")
	assert.NotEmpty(t, RefreshToken.Token, "RefreshToken Token is empty")
	assert.NotEmpty(t, RefreshToken.RefreshToken, "RefreshToken RefreshToken is empty")
}

func TestRegisterDelivery(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controller := MakeAuthController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authRegister := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	deliveryObj := &types.Delivery{
		Name: "test register",
	}

	delivery, err := controller.RegisterDelivery(authRegister, deliveryObj)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, delivery, "Delivery is empty")
}

func MakeAuthController() *AuthController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryClient := repositories.NewClientRepository(database)
	repositoryDelivery := repositories.NewDeliveryRepository(database)
	controller := NewAuthController(
		business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness),
		auth.NewLoginHandler(repositoryAuth, repositoryBusiness),
		auth.NewRefreshTokenHandler(repositoryAuth),
		client.NewRegisterClientHandler(repositoryAuth, repositoryClient),
		delivery.NewRegisterDeliveryHandler(repositoryAuth, repositoryDelivery),
	)
	return controller
}
