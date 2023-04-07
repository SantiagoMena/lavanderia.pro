package client

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestUpdateClientProfileHandler(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	updateHandler := MakeUpdateHandlerToTest()
	registerHandler := MakeClientRegisterHandlerToTestUpdateClient()

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

	clientRegistered, errRegister := registerHandler.Handle(auth, clientObj)

	assert.Nil(t, errRegister, "Error on register client")
	assert.NotEmpty(t, clientRegistered, "Client is empty on register")

	clientRegistered.Name = "UPDATED"
	clientUpdated, errUpdate := updateHandler.Handle(&clientRegistered)

	assert.Nil(t, errUpdate, "Error on update client")
	assert.NotEmpty(t, clientRegistered, "Client is empty on update")
	assert.Equal(t, "UPDATED", clientUpdated.Name, "Client name updated not properly")

}

func MakeUpdateHandlerToTest() *UpdateClientProfileHandler {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	repository := repositories.NewClientRepository(mongo)
	handler := NewUpdateClientProfileHandler(repository)

	return handler
}

func MakeClientRegisterHandlerToTestUpdateClient() *RegisterClientHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryClient := repositories.NewClientRepository(database)
	handler := NewRegisterClientHandler(repositoryAuth, repositoryClient)

	return handler
}
