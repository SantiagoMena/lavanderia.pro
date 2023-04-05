package repositories

import (
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

func TestCreateProduct(t *testing.T) {
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

	business, err := NewProductRepository(mongo).Create(&types.Product{
		Name:  "test",
		Price: 0.123,
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, business, "FindAllBusiness() returns nil result")
	assert.NotEmpty(t, business.CreatedAt, "CreatedAt is empty")
}
