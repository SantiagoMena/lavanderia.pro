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
	createHandler := MakeRegisterToLoginHandler()
	loginHandler := MakeLoginHandler()

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
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
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

func MakeRegisterToLoginHandler() *business.RegisterBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	handler := business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)

	return handler
}

func MakeLoginHandler() *LoginHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	handler := NewLoginHandler(repositoryAuth, repositoryBusiness)

	return handler
}
