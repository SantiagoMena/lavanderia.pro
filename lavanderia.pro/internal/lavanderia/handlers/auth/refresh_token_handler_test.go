package auth

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"strings"
	"testing"
	"time"

	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/repositories"
)

func TestRefreshTokenHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeRegisterToLoginRefresh()
	loginHandler := MakeLoginToRefreshHandler()

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
	assert.NotEmpty(t, businessLogin, "Login Empty")
	assert.NotEmpty(t, businessLogin.Token, "Login Token Empty")

}

func MakeRegisterToLoginRefresh() *business.RegisterBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database)
	handler := business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)

	return handler
}

func MakeLoginToRefreshHandler() *LoginHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database)
	handler := NewLoginHandler(repositoryAuth, repositoryBusiness)

	return handler
}

func MakeRefreshTokenHandler() *RefreshTokenHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database)
	handler := NewRefreshTokenHandler(repositoryAuth, repositoryBusiness)

	return handler
}
