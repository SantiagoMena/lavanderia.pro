package controllers

import (
	"fmt"

	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/handlers/client"
	"lavanderia.pro/internal/lavanderia/handlers/delivery"
	"lavanderia.pro/internal/lavanderia/handlers/product"
	"lavanderia.pro/internal/lavanderia/repositories"
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
	assert.NotEmpty(t, productCreated, "Product is empty")
	assert.NotEmpty(t, productCreated.ID, "Product ID is empty")
	assert.NotEmpty(t, productCreated.CreatedAt, "Product CreatedAt is empty")
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
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	}

	authRegister, errAuth := authController.RegisterBusiness(auth, businessObj)
	assert.Nil(t, errAuth, "Error Auth returns not nil")
	assert.NotEmpty(t, authRegister, "auth is empty")

	productCreated, err := productController.PostProduct(&types.Product{
		Business: authRegister.ID,
		Name:     "test",
		Price:    0.123,
	})

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, productCreated, "productCreated is empty")
	assert.NotEmpty(t, productCreated.ID, "productCreated ID is empty")
	assert.NotEmpty(t, productCreated.CreatedAt, "productCreated CreatedAt is empty")

	productsFound, errFind := productController.GetAllProductsByBusiness(authRegister.ID)

	assert.Nil(t, errFind, "errFind returns not nil")
	assert.NotEmpty(t, productsFound, "productsFound is empty")
}

func TestUpdateProduct(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controller := MakeProductController()

	productCreated, err := controller.PostProduct(&types.Product{
		Name:  "test",
		Price: 0.123,
	})

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, productCreated, "Product is empty")
	assert.NotEmpty(t, productCreated.ID, "Product ID is empty")
	assert.NotEmpty(t, productCreated.CreatedAt, "Product CreatedAt is empty")

	productCreated.Price = 0.789
	productCreated.Name = "UPDATED"
	productUpdated, errUpdate := controller.UpdateProduct(&productCreated)

	assert.Nil(t, errUpdate, "Error returns not nil")
	assert.NotEmpty(t, productUpdated, "Product is empty")
	assert.NotEmpty(t, productUpdated.ID, "Product ID is empty")
	assert.NotEmpty(t, productUpdated.UpdatedAt, "Product UpdatedAt is empty")
	assert.Equal(t, 0.789, productUpdated.Price, "Product Price is not updated")
	assert.Equal(t, "UPDATED", productUpdated.Name, "Product Name is not updated")
}

func MakeProductController() *ProductController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewProductRepository(database)
	controller := NewProductController(
		product.NewCreateProductHandler(repository),
		product.NewGetAllProductsByBusinessHandler(repository),
		product.NewDeleteProductHandler(repository),
		product.NewGetProductHandler(repository),
		product.NewUpdateProductHandler(repository),
	)

	return controller
}

func MakeAuthForProductController() *AuthController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repositoryAuth := repositories.NewAuthRepository(database, config)
	repositoryBusiness := repositories.NewBusinessRepository(database)
	repositoryClient := repositories.NewClientRepository(database)
	repositoryDelivery := repositories.NewDeliveryRepository(database)
	RegisterBusinessHandler := business.NewRegisterBusinessHandler(repositoryAuth, repositoryBusiness)
	RegisterClientHandler := client.NewRegisterClientHandler(repositoryAuth, repositoryClient)
	RegisterDeliveryHandler := delivery.NewRegisterDeliveryHandler(repositoryAuth, repositoryDelivery)
	LoginHandler := auth.NewLoginHandler(repositoryAuth, repositoryBusiness)
	RefreshTokenHandler := auth.NewRefreshTokenHandler(repositoryAuth)
	ChangePasswordHandler := auth.NewChangePasswordHandler(repositoryAuth, repositoryBusiness)

	controller := NewAuthController(
		RegisterBusinessHandler,
		LoginHandler,
		RefreshTokenHandler,
		RegisterClientHandler,
		RegisterDeliveryHandler,
		ChangePasswordHandler,
	)

	return controller
}

func MakeBusinessForProductController() *BusinessController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	businessRepository := repositories.NewBusinessRepository(database)
	authRepository := repositories.NewAuthRepository(database, config)
	deliveryRepository := repositories.NewDeliveryRepository(database)

	controller := NewBusinessController(
		business.NewGetAllBusinessHandler(businessRepository),
		business.NewCreateBusinessHandler(businessRepository),
		business.NewDeleteBusinessHandler(businessRepository),
		business.NewUpdateBusinessHandler(businessRepository),
		business.NewGetBusinessHandler(businessRepository),
		business.NewGetAllBusinessByAuthHandler(businessRepository),
		business.NewRegisterBusinessDeliveryHandler(authRepository, businessRepository, deliveryRepository),
	)

	return controller
}
