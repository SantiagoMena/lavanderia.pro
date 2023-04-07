package client

import (
	"github.com/joho/godotenv"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"

	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestPostClient(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	handler := MakePostClientHandlerToTest()

	createHandler := MakeClientRegisterHandlerToTestPostClient()

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

	clientPosted, errPost := handler.Handle(&clientRegistered)

	assert.NotNil(t, errPost, "Error on post client")
	assert.Empty(t, clientPosted, "Client Posted Empty")
	assert.Equal(t, "test register", clientPosted.Name, "Client Name nil")
}

func MakePostClientHandlerToTest() *PostClientHandler {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	repository := repositories.NewClientRepository(mongo)
	handler := NewPostClientHandler(repository)

	return handler
}

func MakeClientRegisterHandlerToTestPostClient() *RegisterClientHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryClient := repositories.NewClientRepository(database)
	handler := NewRegisterClientHandler(repositoryAuth, repositoryClient)

	return handler
}
