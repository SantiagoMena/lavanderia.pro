package auth

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestLoginHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeRegisterToLoginBusinessHandler()
	loginHandler := MakeLoginBusinessHandler()

	pwd := []byte("PwD")
	// password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	// assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(pwd),
	}

	businessObj := &types.Business{
		Name: "test register",
		Lat:  0.321,
		Long: 0.321,
	}

	businessRegistered, errRegister := createHandler.Handle(auth, businessObj)

	assert.Nil(t, errRegister, "Error on register business")
	assert.NotEmpty(t, businessRegistered, "Business is empty on register")

	businessLogin, errLogin := loginHandler.Handle(&types.Auth{
		Email:    auth.Email,
		Password: string(pwd),
	})
	assert.Nil(t, errLogin, "Error on login business")
	// assert.Equal(t, errLogin, errors.New("auth already exists"), "Error auth already exists not work")
	assert.NotEmpty(t, businessLogin, "Business Login Empty")

}

func MakeRegisterToLoginBusinessHandler() *business.RegisterBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database)
	handler := business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)

	return handler
}

func MakeLoginBusinessHandler() *LoginBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database)
	handler := NewLoginBusinessHandler(repositoryAuth, repositoryBusiness)

	return handler
}
