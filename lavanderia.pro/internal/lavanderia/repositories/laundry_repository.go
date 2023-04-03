package repositories

import (
	"context"

	"lavanderia.pro/api/types"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/internal/lavanderia/databases"
)

var collection = "laundry"

type LaundryRepository struct {
	database databases.Database
}

func NewLaundryRepository(database databases.Database) *LaundryRepository {
	return &LaundryRepository{
		database: database,
	}
}

func (laundryRepository *LaundryRepository) FindAllLaundries() ([]types.Laundry, error) {
	laundries := []types.Laundry{}

	laundriesDb, err := laundryRepository.database.FindAll(collection)

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

func (laundryRepository *LaundryRepository) Create(laundry *types.Laundry) (types.Laundry, error) {

	laundryDb, err := laundryRepository.database.Create("laundry", bson.D{
		{Key: "name", Value: laundry.Name},
		{Key: "lat", Value: laundry.Lat},
		{Key: "long", Value: laundry.Long},
	})

	if err != nil {
		return types.Laundry{}, err
	}

	insertedId := laundryDb.InsertedID.(primitive.ObjectID).Hex()

	newLaundry := types.Laundry{
		ID:   insertedId,
		Name: laundry.Name,
		Lat:  laundry.Lat,
		Long: laundry.Long,
	}

	return newLaundry, nil
}
