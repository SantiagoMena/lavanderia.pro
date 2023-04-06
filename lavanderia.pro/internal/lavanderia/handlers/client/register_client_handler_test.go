package client

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"errors"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestRegisterHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeClientRegisterHandlerToTest()

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

	clientRegisteredSameEmail, errRegisterSame := createHandler.Handle(auth, clientObj)
	assert.Equal(t, errRegisterSame, errors.New("auth already exists"), "Error auth already exists not work")
	assert.Equal(t, types.Client{}, clientRegisteredSameEmail, "Error auth already exists return empty client")

}

func MakeClientRegisterHandlerToTest() *RegisterClientHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryClient := repositories.NewClientRepository(database)
	handler := NewRegisterClientHandler(repositoryAuth, repositoryClient)

	return handler
}
