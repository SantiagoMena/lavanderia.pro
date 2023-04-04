package business

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
	createHandler := MakeCreateBusinessToRegisterHandler()
	// updateHandler := MakeUpdateBusinessHandler()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	businessObj := &types.Business{
		Name: "test register",
		Lat:  0.321,
		Long: 0.321,
	}

	businessRegistered, errRegister := createHandler.Handle(auth, businessObj)

	assert.Nil(t, errRegister, "Error on register business")
	assert.NotEmpty(t, businessRegistered, "Business is empty on register")

	businessRegisteredSameEmail, errRegisterSame := createHandler.Handle(auth, businessObj)
	assert.Equal(t, errRegisterSame, errors.New("auth already exists"), "Error auth already exists not work")
	assert.Equal(t, types.Business{}, businessRegisteredSameEmail, "Error auth already exists return empty business")

}

func MakeUpdateBusinessHandler() *UpdateBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewBusinessRepository(database)
	handler := NewUpdateBusinessHandler(repository)

	return handler
}

func MakeCreateBusinessToRegisterHandler() *RegisterBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	handler := NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)

	return handler
}
