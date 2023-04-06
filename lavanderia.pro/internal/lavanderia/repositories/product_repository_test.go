package repositories

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
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

func TestGetAllProductsByBusiness(t *testing.T) {
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

func TestDeleteProduct(t *testing.T) {
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

	product, err := NewProductRepository(mongo).Create(&types.Product{
		Name:  "test",
		Price: 0.123,
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, product, "Create() returns nil result")
	assert.NotEmpty(t, product.CreatedAt, "CreatedAt is empty")

	productDeleted, errDelete := NewProductRepository(mongo).Delete(&types.Product{
		ID: product.ID,
	})

	var productDeletedObject types.Product

	// convert m to s
	productDeletedObj, _ := bson.Marshal(productDeleted)
	bson.Unmarshal(productDeletedObj, &productDeletedObject)

	assert.Equal(t, errDelete, nil, "Delete() returns error")
	assert.NotNil(t, productDeleted, "Delete() returns nil result")
	assert.NotEmpty(t, productDeletedObject.DeletedAt, "DeletedAt is empty")

}

func TestUpdateProduct(t *testing.T) {
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

	product, err := NewProductRepository(mongo).Create(&types.Product{
		Name:  "test",
		Price: 0.123,
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, product, "Create() returns nil result")
	assert.NotEmpty(t, product.CreatedAt, "CreatedAt is empty")

	productUpdated, errDelete := NewProductRepository(mongo).Update(&types.Product{
		ID:    product.ID,
		Price: 123,
	})

	var productUpdatedObject types.Product

	// convert m to s
	productUpdatedObj, _ := bson.Marshal(productUpdated)
	bson.Unmarshal(productUpdatedObj, &productUpdatedObject)

	assert.Equal(t, errDelete, nil, "Update() returns error")
	assert.NotNil(t, productUpdated, "Update() returns nil result")
	assert.NotEmpty(t, productUpdatedObject.UpdatedAt, "UpdatedAt is empty")
	assert.Equal(t, float64(123), productUpdatedObject.Price, "price is different from updated")

}
