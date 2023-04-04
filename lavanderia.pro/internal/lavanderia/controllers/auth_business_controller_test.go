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
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestRegisterBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controller := MakeAuthBusinessController()

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
		Lat:  0.321,
		Long: 0.321,
	}

	business, err := controller.RegisterBusiness(authRegister, businessObj)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, business, "Business is empty")
}

func TestLogin(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controller := MakeAuthBusinessController()

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
		Lat:  0.321,
		Long: 0.321,
	}

	business, err := controller.RegisterBusiness(authLogin, businessObj)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, business, "Business is empty")

	businessLogged, errLogin := controller.Login(authLogin)

	assert.Nil(t, errLogin, "Error Login returns not nil")
	assert.NotEmpty(t, businessLogged, "Business Logged is empty")
}

func MakeAuthBusinessController() *AuthBusinessController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	controller := NewAuthBusinessController(
		business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness),
		auth.NewLoginHandler(repositoryAuth, repositoryBusiness),
	)
	return controller
}
