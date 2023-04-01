package repositories

import (
	"context"
	"log"

	"lavanderia.pro/api/types"

	"lavanderia.pro/internal/lavanderia/databases"
)

var collection = "laundry"

func FindAllLaundries() []types.Laundry {
	laundries := []types.Laundry{}

	laundriesDb, err := databases.FindAll(collection)

	if err != nil {
		log.Panic(err)
	}

	log.Panic(laundriesDb)

	for laundriesDb.Next(context.TODO()) {
		var laundry types.Laundry

		if err := laundriesDb.Decode(&laundry); err != nil {
			log.Panic(err)

		}

		laundries = append(laundries, laundry)
	}

	return laundries
}
