package business

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"lavanderia.pro/internal/lavanderia/config"

	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
)

func TestGetAllBusinessByAuthHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	// Register
	registerBusinessHandler := MakeRegisterBusinessToTestGetAllByAuthHandler()
	loginHandler := MakeLoginToTestGetAllByAuthHandler()

	pwd := []byte("PwD")

	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(pwd),
	}

	businessObj := &types.Business{
		Name: "test register",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	businessRegistered, errRegister := registerBusinessHandler.Handle(auth, businessObj)
	assert.Nil(t, errRegister, "errRegister returns not nil")
	assert.NotEmpty(t, businessRegistered, "BusinessRegistered empty")

	// Auth
	businessLogged, errLogin := loginHandler.Handle(auth)
	assert.Nil(t, errLogin, "errLogin returns not nil")

	// Test
	handler := MakeGetAllByAuthBusinessHandler()

	allBusiness, err := handler.Handle(businessLogged.Token)

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, allBusiness, "Business is empty")
}

func MakeGetAllByAuthBusinessHandler() *GetAllBusinessByAuthHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewBusinessRepository(database)
	handler := NewGetAllBusinessByAuthHandler(repository)

	return handler
}

func MakeRegisterBusinessToTestGetAllByAuthHandler() *RegisterBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	registerBusinessHandler := NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)

	return registerBusinessHandler
}

func MakeLoginToTestGetAllByAuthHandler() *auth.LoginHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	loginHandler := auth.NewLoginHandler(repositoryAuth, repositoryBusiness)

	return loginHandler
}
