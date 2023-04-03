package repositories

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	"testing"
)

func TestFindAllLaundries(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	laundries, err := NewLaundryRepository(mongo).FindAllLaundries()

	laundriesExpect := []types.Laundry{}

	mongo2 := databases.NewMongoDatabase(config)
	laundriesDb, err := mongo2.FindAll("laundry")

	if err != nil {
		fmt.Println(err)
	}

	for laundriesDb.Next(context.TODO()) {
		var laundry types.Laundry

		if err := laundriesDb.Decode(&laundry); err != nil {
			fmt.Println(err)
		}

		laundriesExpect = append(laundriesExpect, laundry)
	}

	assert.Equal(t, err, nil, "FindAllLAundries() returns error")
	assert.NotNil(t, laundries, laundriesExpect, "FindAllLAundries() returns nil result")
}

func TestCreateLaundry(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	laundry, err := NewLaundryRepository(mongo).Create(&types.Laundry{
		Name: "test",
		Lat:  0.123,
		Long: 0.123,
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, laundry, "FindAllLAundries() returns nil result")
	assert.NotEmpty(t, laundry.CreatedAt, "CreatedAt is empty")
}

func TestDeleteLaundry(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	laundry, err := NewLaundryRepository(mongo).Create(&types.Laundry{
		Name: "test",
		Lat:  0.321,
		Long: 0.321,
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, laundry, "Create() returns nil result")
	assert.NotEmpty(t, laundry.CreatedAt, "CreatedAt is empty")
	assert.NotEmpty(t, laundry.ID, "ID is empty")

	laundryDeleted, errDelete := NewLaundryRepository(mongo).Delete(&laundry)

	assert.Equal(t, errDelete, nil, "Delete() returns error")
	assert.NotNil(t, laundryDeleted, "Delete() returns nil result")
	assert.NotEmpty(t, laundryDeleted.CreatedAt, "CreatedAt is empty")
	assert.NotEmpty(t, laundryDeleted.DeletedAt, "DeletedAt is empty")
	assert.NotEmpty(t, laundryDeleted.ID, "ID is empty")

}

func TestGetLaundry(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	config := config.NewConfig()

	mongo := databases.NewMongoDatabase(config)

	laundry, err := NewLaundryRepository(mongo).Create(&types.Laundry{
		Name: "test",
		Lat:  0.321,
		Long: 0.321,
	})

	assert.Equal(t, err, nil, "Create() returns error")
	assert.NotNil(t, laundry, "Create() returns nil result")
	assert.NotEmpty(t, laundry.CreatedAt, "CreatedAt is empty")
	assert.NotEmpty(t, laundry.ID, "ID is empty")

	laundryGotten, errDelete := NewLaundryRepository(mongo).Get(&laundry)

	assert.Equal(t, errDelete, nil, "Delete() returns error")
	assert.NotNil(t, laundryGotten, "Delete() returns nil result")
	assert.NotEmpty(t, laundryGotten.CreatedAt, "CreatedAt is empty")
	assert.NotEmpty(t, laundryGotten.ID, "ID is empty")

}
