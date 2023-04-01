package repositories

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/databases"
	"testing"
)

func TestFindAllLaundries(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	laundries, err := FindAllLaundries()

	laundriesExpect := []types.Laundry{}

	laundriesDb, err := databases.FindAll(collection)

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
	assert.Equal(t, laundries, laundriesExpect, "FindAllLAundries() returns different result")
}
