package controllers

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/repositories"
	"testing"
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
		Lat:  0.123,
		Long: 0.123,
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
		Lat:  0.123,
		Long: 0.123,
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
		Lat:  0.123,
		Long: 0.123,
	})

	assert.Nil(t, err, "Error returns not nil on create business to delete")
	assert.NotEmpty(t, business, "Business is empty on create business to delete")
	assert.NotEmpty(t, business.ID, "Business ID is empty on create business to delete")
	assert.Equal(t, "test", business.Name, "Name not saved properly")
	assert.Equal(t, 0.123, business.Lat, "Lat not saved properly")
	assert.Equal(t, 0.123, business.Long, "Long not saved properly")

	businessUpdated, errUpdate := controller.UpdateBusiness(&types.Business{
		ID:   business.ID,
		Name: "updated",
		Lat:  0.321,
		Long: 0.321,
	})
	assert.Nil(t, errUpdate, "Error returns not nil on delete business")
	assert.NotEmpty(t, businessUpdated, "Business is empty on delete business")
	assert.NotEmpty(t, businessUpdated.ID, "Business ID is empty on delete business")
	assert.NotEmpty(t, businessUpdated.UpdatedAt, "Business UpdatedAt is empty on delete business")
	assert.Equal(t, "updated", businessUpdated.Name, "Name not save properly")
	assert.Equal(t, 0.321, businessUpdated.Lat, "Lat not save properly")
	assert.Equal(t, 0.321, businessUpdated.Long, "Long not save properly")
}

func TestGetBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}
	controller := MakeController()

	business, err := controller.PostBusiness(&types.Business{
		Name: "test",
		Lat:  0.123,
		Long: 0.123,
	})

	assert.Nil(t, err, "Error returns not nil on create business to delete")
	assert.NotEmpty(t, business, "Business is empty on create business to delete")
	assert.NotEmpty(t, business.ID, "Business ID is empty on create business to delete")
	assert.Equal(t, "test", business.Name, "Name not saved properly")
	assert.Equal(t, 0.123, business.Lat, "Lat not saved properly")
	assert.Equal(t, 0.123, business.Long, "Long not saved properly")

	businessGotten, errGet := controller.GetBusiness(&business)
	assert.Nil(t, errGet, "Error returns not nil on delete business")
	assert.NotEmpty(t, businessGotten, "Business is empty on delete business")
	assert.NotEmpty(t, businessGotten.ID, "Business ID is empty on delete business")
	assert.NotEmpty(t, businessGotten.CreatedAt, "Business CreatedAt is empty on delete business")
	assert.Equal(t, "test", businessGotten.Name, "Name not get properly")
	assert.Equal(t, 0.123, businessGotten.Lat, "Lat not get properly")
	assert.Equal(t, 0.123, businessGotten.Long, "Long not get properly")
}

func MakeController() *BusinessController {
	config := config.NewConfig()
	database := databases.NewMongoDatabase(config)
	repository := repositories.NewBusinessRepository(database)
	controller := NewBusinessController(
		business.NewGetAllBusinessHandler(repository),
		business.NewCreateBusinessHandler(repository),
		business.NewDeleteBusinessHandler(repository),
		business.NewUpdateBusinessHandler(repository),
		business.NewGetBusinessHandler(repository),
		business.NewGetAllBusinessByAuthHandler(repository),
	)
	return controller
}
