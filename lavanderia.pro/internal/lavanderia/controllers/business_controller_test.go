package controllers

import (
	"fmt"

	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/handlers/client"
	"lavanderia.pro/internal/lavanderia/handlers/delivery"
	"lavanderia.pro/internal/lavanderia/repositories"
)

func TestGetAllBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controller := MakeController()
	business, err := controller.GetAllBusiness()

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, business, "Business is empty")
}

func TestPostBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controller := MakeController()

	business, err := controller.PostBusiness(&types.Business{
		Name: "test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})

	assert.Nil(t, err, "Error returns not nil")
	assert.NotEmpty(t, business, "Business is empty")
	assert.NotEmpty(t, business.ID, "Business ID is empty")
}

func TestDeleteBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controller := MakeController()

	business, err := controller.PostBusiness(&types.Business{
		Name: "test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})

	assert.Nil(t, err, "Error returns not nil on create business to delete")
	assert.NotEmpty(t, business, "Business is empty on create business to delete")
	assert.NotEmpty(t, business.ID, "Business ID is empty on create business to delete")

	businessDeleted, errDelete := controller.DeleteBusiness(&business)
	assert.Nil(t, errDelete, "Error returns not nil on delete business")
	assert.NotEmpty(t, businessDeleted, "Business is empty on delete business")
	assert.NotEmpty(t, businessDeleted.ID, "Business ID is empty on delete business")
	assert.NotEmpty(t, businessDeleted.DeletedAt, "Business DeletedAt is empty on delete business")
}

func TestUpdateBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controller := MakeController()

	business, err := controller.PostBusiness(&types.Business{
		Name: "test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})

	assert.Nil(t, err, "Error returns not nil on create business to delete")
	assert.NotEmpty(t, business, "Business is empty on create business to delete")
	assert.NotEmpty(t, business.ID, "Business ID is empty on create business to delete")
	assert.Equal(t, "test", business.Name, "Name not saved properly")
	assert.NotEmpty(t, business.Position, "Position not saved properly")

	businessUpdated, errUpdate := controller.UpdateBusiness(&types.Business{
		ID:   business.ID,
		Name: "updated",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})
	assert.Nil(t, errUpdate, "Error returns not nil on delete business")
	assert.NotEmpty(t, businessUpdated, "Business is empty on delete business")
	assert.NotEmpty(t, businessUpdated.ID, "Business ID is empty on delete business")
	assert.NotEmpty(t, businessUpdated.UpdatedAt, "Business UpdatedAt is empty on delete business")
	assert.Equal(t, "updated", businessUpdated.Name, "Name not save properly")
	assert.NotEmpty(t, businessUpdated.Position, "Position not save properly")
}

func TestGetBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controller := MakeController()

	business, err := controller.PostBusiness(&types.Business{
		Name: "test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})

	assert.Nil(t, err, "Error returns not nil on create business to delete")
	assert.NotEmpty(t, business, "Business is empty on create business to delete")
	assert.NotEmpty(t, business.ID, "Business ID is empty on create business to delete")
	assert.Equal(t, "test", business.Name, "Name not saved properly")
	assert.NotEmpty(t, business.Position, "Position not saved properly")

	businessGotten, errGet := controller.GetBusiness(&business)
	assert.Nil(t, errGet, "Error returns not nil on delete business")
	assert.NotEmpty(t, businessGotten, "Business is empty on delete business")
	assert.NotEmpty(t, businessGotten.ID, "Business ID is empty on delete business")
	assert.NotEmpty(t, businessGotten.CreatedAt, "Business CreatedAt is empty on delete business")
	assert.Equal(t, "test", businessGotten.Name, "Name not get properly")
	assert.NotEmpty(t, businessGotten.Position, "Position not saved properly")
}

func TestRegisterBusinessDelivery(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controller := MakeController()

	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")
	ti := time.Now()
	email := []string{"new@", ti.String(), "test.com"}

	authRegister := &types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	}

	deliveryCreated, err := controller.RegisterBusinessDelivery(
		authRegister,
		&types.Business{
			Name: "test",
			Position: types.Geometry{
				Type:        "Point",
				Coordinates: []float64{-71.327767, -41.138444},
			},
		},
		&types.Delivery{
			Name: "test_delivery",
		},
	)

	assert.Nil(t, err, "Error returns not nil on create business delivery")
	assert.NotEmpty(t, deliveryCreated, "deliveryCreated is empty on create business delivery")
	assert.NotEmpty(t, deliveryCreated.ID, "deliveryCreated ID is empty on create business delivery")
	assert.Equal(t, "test_delivery", deliveryCreated.Name, "Name not saved properly")

}

func MakeController() *BusinessController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewBusinessRepository(database)
	deliveryRepository := repositories.NewDeliveryRepository(database)
	authRepository := repositories.NewAuthRepository(database, config)

	controller := NewBusinessController(
		business.NewGetAllBusinessHandler(repository),
		business.NewCreateBusinessHandler(repository),
		business.NewDeleteBusinessHandler(repository),
		business.NewUpdateBusinessHandler(repository),
		business.NewGetBusinessHandler(repository),
		business.NewGetAllBusinessByAuthHandler(repository),
		business.NewRegisterBusinessDeliveryHandler(authRepository, repository, deliveryRepository),
		business.NewSearchBusinessHandler(repository),
	)
	return controller
}

func TestGetAllBusinessByAuth(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	controller := MakeController()
	controllerAuth := MakeAuthForBusinessController()

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

	businessRegistered, errRegister := controllerAuth.RegisterBusiness(auth, businessObj)
	assert.Nil(t, errRegister, "Error returns not nil on register")
	assert.NotEmpty(t, businessRegistered, "businessRegistered is empty on register")

	businessLogged, errLogin := controllerAuth.Login(auth)
	assert.Nil(t, errLogin, "Error returns not nil on login")
	assert.NotEmpty(t, businessLogged, "businessLogged is empty on register")

	business, err := controller.GetAllBusinessByAuth(businessLogged.Token)

	assert.Nil(t, err, "Error returns not nil on create business to delete")
	assert.NotEmpty(t, business, "Business is empty on create business to delete")
}

// func MakeController() *BusinessController {
// 	config := config.NewConfig()
// 	database := databases.NewMongoDatabase(config)
// 	repository := repositories.NewBusinessRepository(database)
// 	controller := NewBusinessController(
// 		business.NewGetAllBusinessHandler(repository),
// 		business.NewCreateBusinessHandler(repository),
// 		business.NewDeleteBusinessHandler(repository),
// 		business.NewUpdateBusinessHandler(repository),
// 		business.NewGetBusinessHandler(repository),
// 		business.NewGetAllBusinessByAuthHandler(repository),
// 	)
// 	return controller
// }

func MakeAuthForBusinessController() *AuthController {
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
