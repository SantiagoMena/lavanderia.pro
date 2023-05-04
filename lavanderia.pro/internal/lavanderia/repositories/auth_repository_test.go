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
	"strings"
	"testing"
	"time"
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

	auth, err := NewAuthRepository(mongo, config).Create(&types.Auth{
		Email:    "new@test.com",
		Password: string(password),
	})

	mongo2 := databases.NewMongoDatabase(config)
	authFound, errFind := NewAuthRepository(mongo2, config).GetByEmail(&types.Auth{
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

	auth, err := NewAuthRepository(mongo, config).Create(&types.Auth{
		Email:    "new@test.com",
		Password: string(password),
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, auth, "FindAllBusiness() returns nil result")
	assert.NotEmpty(t, auth.CreatedAt, "CreatedAt is empty")
}

func TestCreateJWT(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")

	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	auth, err := NewAuthRepository(mongo, config).Create(&types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, auth, "FindAllBusiness() returns nil result")
	assert.NotEmpty(t, auth.CreatedAt, "CreatedAt is empty")

	jwt, errJWT := NewAuthRepository(mongo, config).CreateJWT(&auth)
	assert.Equal(t, errJWT, nil, "Create() returns error")
	assert.NotNil(t, jwt, "Login() returns nil result")
	assert.NotEmpty(t, jwt.Token, "Token is empty")
}

func TestUpdatePassword(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)
	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")

	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}
	auth, err := NewAuthRepository(mongo, config).Create(&types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	})

	mongo3 := databases.NewMongoDatabase(config)

	authUpdated, errAuthUpdate := NewAuthRepository(mongo3, config).UpdatePassword(&types.Auth{
		ID:       auth.ID,
		Password: "NEW_FB_ID",
	})

	assert.NotNil(t, authUpdated, "FindAllBusiness() returns nil result")
	assert.Nil(t, errAuthUpdate, "errAuthUpdate error")

	mongo2 := databases.NewMongoDatabase(config)
	authFound, errFind := NewAuthRepository(mongo2, config).GetByEmail(&types.Auth{
		Email: "new@test.com",
	})

	assert.Equal(t, errFind, nil, "GetByEmail() returns error")
	assert.Equal(t, err, nil, "Create() returns error")

	assert.NotEqual(t, auth.Password, authFound.Password, "Entity not updated password")
}
