package repositories

import (
	"context"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
)

func TestFindAllBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	business, errFindBusiness := NewBusinessRepository(mongo).FindAllBusiness()
	assert.Nil(t, errFindBusiness, "Error on FindAllBusiness()")

	businessExpect := []types.Business{}

	mongo2 := databases.NewMongoDatabase(config)
	businessDb, errFindAll := mongo2.FindAll("business")
	assert.Nil(t, errFindAll, "FindAllBusiness() returns error")

	for businessDb.Next(context.TODO()) {
		var business types.Business

		if err := businessDb.Decode(&business); err != nil {
			fmt.Println(err)
		}

		businessExpect = append(businessExpect, business)
	}

	assert.NotNil(t, business, businessExpect, "FindAllBusiness() returns nil result")
}

func TestCreateBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	business, err := NewBusinessRepository(mongo).Create(&types.Business{
		Name: "test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, business, "FindAllBusiness() returns nil result")
	assert.NotEmpty(t, business.CreatedAt, "CreatedAt is empty")
}

func TestDeleteBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	business, err := NewBusinessRepository(mongo).Create(&types.Business{
		Name: "test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, business, "Create() returns nil result")
	assert.NotEmpty(t, business.CreatedAt, "CreatedAt is empty")
	assert.NotEmpty(t, business.ID, "ID is empty")

	businessDeleted, errDelete := NewBusinessRepository(mongo).Delete(&business)

	assert.Equal(t, errDelete, nil, "Delete() returns error")
	assert.NotNil(t, businessDeleted, "Delete() returns nil result")
	assert.NotEmpty(t, businessDeleted.CreatedAt, "CreatedAt is empty")
	assert.NotEmpty(t, businessDeleted.DeletedAt, "DeletedAt is empty")
	assert.NotEmpty(t, businessDeleted.ID, "ID is empty")

}

func TestGetBusiness(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	business, err := NewBusinessRepository(mongo).Create(&types.Business{
		Name: "test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, business, "Create() returns nil result")
	assert.NotEmpty(t, business.CreatedAt, "CreatedAt is empty")
	assert.NotEmpty(t, business.ID, "ID is empty")

	businessGotten, errDelete := NewBusinessRepository(mongo).Get(&business)

	assert.Equal(t, errDelete, nil, "Delete() returns error")
	assert.NotNil(t, businessGotten, "Delete() returns nil result")
	assert.NotEmpty(t, businessGotten.CreatedAt, "CreatedAt is empty")
	assert.NotEmpty(t, businessGotten.ID, "ID is empty")

}
