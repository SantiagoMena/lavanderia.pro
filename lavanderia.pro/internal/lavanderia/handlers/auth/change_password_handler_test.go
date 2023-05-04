package auth

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestChangePasswordHandle(t *testing.T) {
	if err := godotenv.Load("../../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	createHandler := MakeRegisterToChangePasswordHandler()
	loginHandler := MakeLoginChangePasswordHandler()
	changePasswordHandler := MakeChangePasswordHandler()
	authRepository := MakeAuthRepositoryToChangePasswordHandler()

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

	businessRegistered, errRegister := createHandler.Handle(auth, businessObj)

	assert.Nil(t, errRegister, "Error on register business")
	assert.NotEmpty(t, businessRegistered, "Business is empty on register")

	businessLogin, errLogin := loginHandler.Handle(&types.Auth{
		Email:    auth.Email,
		Password: string(pwd),
	})

	assert.Nil(t, errLogin, "Error on login business")
	assert.NotEmpty(t, businessLogin, "Business Login Empty")

	authFound, errAuthFound := authRepository.GetByEmail(auth)

	assert.Nil(t, errAuthFound, "Error on login business")
	assert.NotEmpty(t, authFound, "Business Login Empty")

	newPassword := "NewPassword"
	newPwd := []byte(newPassword)

	changePassword, errChangePassword := changePasswordHandler.Handle(authFound.ID, &types.NewPassword{
		ID:          authFound.ID,
		Password:    string(pwd),
		NewPassword: string(newPwd),
	})

	assert.Nil(t, errChangePassword, "Error on change password")
	assert.NotEmpty(t, changePassword, "changePassword Empty")

	authFoundChanged, errAuthFoundChanged := authRepository.GetByEmail(auth)
	assert.NotEmpty(t, authFoundChanged, "authFoundChanged Empty")

	fmt.Println(authFoundChanged.Password)

	assert.Nil(t, errAuthFoundChanged, "Error on getAuth changed password")

	newPwdChanged := []byte(authFoundChanged.Password)
	passwordChanged := bcrypt.CompareHashAndPassword(newPwdChanged, newPwd)

	assert.Nil(t, passwordChanged, "Password not change properly")
}

func MakeRegisterToChangePasswordHandler() *business.RegisterBusinessHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	handler := business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)

	return handler
}

func MakeLoginChangePasswordHandler() *LoginHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	handler := NewLoginHandler(repositoryAuth, repositoryBusiness)

	return handler
}

func MakeChangePasswordHandler() *ChangePasswordHandler {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	handler := NewChangePasswordHandler(repositoryAuth, repositoryBusiness)

	return handler
}

func MakeAuthRepositoryToChangePasswordHandler() *repositories.AuthRepository {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)

	return repositoryAuth
}
