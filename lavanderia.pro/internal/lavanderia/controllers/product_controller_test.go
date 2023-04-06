package controllers

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/handlers/product"
	"lavanderia.pro/internal/lavanderia/repositories"
	"strings"
	"testing"
	"time"
)

func TestPostProduct(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controller := MakeProductController()

	productCreated, err := controller.PostProduct(&types.Product{
		Name:  "test",
		Price: 0.123,
	})

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, productCreated, "Business is empty")
	assert.NotEmpty(t, productCreated.ID, "Business ID is empty")
	assert.NotEmpty(t, productCreated.CreatedAt, "Business CreatedAt is empty")
}

func TestGetAllProductsByBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	productController := MakeProductController()
	// businessController := MakeBusinessForProductController()
	authController := MakeAuthForProductController()

	pwd := []byte("PwD")

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

	authRegister, errAuth := authController.RegisterBusiness(auth, businessObj)
	assert.Nil(t, errAuth, "Error Auth returns not nil")
	assert.NotEmpty(t, authRegister, "auth is empty")

	productCreated, err := productController.PostProduct(&types.Product{
		Business: auth.ID,
		Name:     "test",
		Price:    0.123,
	})

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, productCreated, "productCreated is empty")
	assert.NotEmpty(t, productCreated.ID, "productCreated ID is empty")
	assert.NotEmpty(t, productCreated.CreatedAt, "productCreated CreatedAt is empty")

	productsFound, errFind := productController.GetAllProductsByBusiness(string(auth.ID))

	assert.Nil(t, errFind, "errFind returns not nil")
	assert.NotEmpty(t, productsFound, "productsFound is empty")
}

func MakeProductController() *ProductController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewProductRepository(database)
	controller := NewProductController(
		product.NewCreateProductHandler(repository),
		product.NewGetAllProductsByBusinessHandler(repository),
	)

	return controller
}

func MakeAuthForProductController() *AuthController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	RegisterBusinessHandler := business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)
	LoginHandler := auth.NewLoginHandler(repositoryAuth, repositoryBusiness)
	RefreshTokenHandler := auth.NewRefreshTokenHandler(repositoryAuth)

	controller := NewAuthController(
		RegisterBusinessHandler,
		LoginHandler,
		RefreshTokenHandler,
	)

	return controller
}

func MakeBusinessForProductController() *BusinessController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	// repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	// RegisterBusinessHandler := business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)
	// LoginHandler := auth.NewLoginHandler(repositoryAuth, repositoryBusiness)
	// RefreshTokenHandler := auth.NewRefreshTokenHandler(repositoryAuth)

	controller := NewBusinessController(
		business.NewGetAllBusinessHandler(repositoryBusiness),
		business.NewCreateBusinessHandler(repositoryBusiness),
		business.NewDeleteBusinessHandler(repositoryBusiness),
		business.NewUpdateBusinessHandler(repositoryBusiness),
		business.NewGetBusinessHandler(repositoryBusiness),
		business.NewGetAllBusinessByAuthHandler(repositoryBusiness),
	)

	return controller
}
