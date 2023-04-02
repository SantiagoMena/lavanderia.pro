package repositories

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
	// "os"
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
