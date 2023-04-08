package client

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"strings"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
)

func TestGetClient(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeClientRegisterHandlerToTestGet()
	getHandler := MakeGetClientHandlerToTest()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	clientObj := &types.Client{
		Name: "test register",
	}

	clientRegistered, errRegister := createHandler.Handle(auth, clientObj)

	assert.Nil(t, errRegister, "Error on register client")
	assert.NotEmpty(t, clientRegistered, "Client is empty on register")

	clientFound, errFind := getHandler.Handle(&clientRegistered)

	assert.Nil(t, errFind, "Error on getClient()")
	assert.NotEmpty(t, clientFound, "Client is empty on getClient()")

}

func MakeClientRegisterHandlerToTestGet() *RegisterClientHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryClient := repositories.NewClientRepository(database)
	handler := NewRegisterClientHandler(repositoryAuth, repositoryClient)

	return handler
}

func MakeGetClientHandlerToTest() *GetClientHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryClient := repositories.NewClientRepository(database)
	handler := NewGetClientHandler(repositoryAuth, repositoryClient)

	return handler
}
