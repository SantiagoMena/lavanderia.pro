package repositories

import (
	"context"

	"lavanderia.pro/api/types"

	"lavanderia.pro/internal/lavanderia/databases"
)

var collection = "laundry"

func FindAllLaundries() ([]types.Laundry, error) {
	laundries := []types.Laundry{}

	laundriesDb, err := databases.FindAll(collection)

	if err != nil {
		return nil, err
	}

	for laundriesDb.Next(context.TODO()) {
		var laundry types.Laundry

		if err := laundriesDb.Decode(&laundry); err != nil {
			return nil, err
		}

		laundries = append(laundries, laundry)
	}

	return laundries, nil
}
