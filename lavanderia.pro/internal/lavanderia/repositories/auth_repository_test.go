package repositories

import (
	// "context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"testing"
)

func TestFindByEmail(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)
	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")

	auth, err := NewAuthRepository(mongo).Create(&types.Auth{
		Email:    "new@test.com",
		Password: password,
	})

	mongo2 := databases.NewMongoDatabase(config)
	authFound, errFind := NewAuthRepository(mongo2).GetByEmail(&types.Auth{
		Email: "new@test.com",
	})

	assert.Equal(t, errFind, nil, "GetByEmail() returns error")
	assert.Equal(t, err, nil, "Create() returns error")
	assert.Equal(t, auth.Email, authFound.Email, "Emails are not equal")
	assert.NotNil(t, auth, authFound, "FindAllBusiness() returns nil result")
}

func TestCreateAuth(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")

	auth, err := NewAuthRepository(mongo).Create(&types.Auth{
		Email:    "new@test.com",
		Password: password,
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, auth, "FindAllBusiness() returns nil result")
	assert.NotEmpty(t, auth.CreatedAt, "CreatedAt is empty")
}
