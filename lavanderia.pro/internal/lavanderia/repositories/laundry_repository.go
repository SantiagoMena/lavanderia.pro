package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/api/types"
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
	t := time.Now()
	laundry.CreatedAt = &t

	laundryDb, err := laundryRepository.database.Create("laundry", bson.D{
		{Key: "name", Value: laundry.Name},
		{Key: "lat", Value: laundry.Lat},
		{Key: "long", Value: laundry.Long},
		{Key: "created_at", Value: laundry.CreatedAt},
	})

	if err != nil {
		return types.Laundry{}, err
	}

	insertedId := laundryDb.InsertedID.(primitive.ObjectID).Hex()

	newLaundry := types.Laundry{
		ID:        insertedId,
		Name:      laundry.Name,
		Lat:       laundry.Lat,
		Long:      laundry.Long,
		CreatedAt: laundry.CreatedAt,
	}

	return newLaundry, nil
}

func (laundryRepository *LaundryRepository) Delete(laundry *types.Laundry) (types.Laundry, error) {
	t := time.Now()
	laundry.DeletedAt = &t

	id, _ := primitive.ObjectIDFromHex(laundry.ID)

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"deleted_at", laundry.DeletedAt}}}}

	objectUpdated, err := laundryRepository.database.UpdateOne(collection, filter, update)
	if err != nil {
		panic(err)
	}

	if err != nil {
		return types.Laundry{}, err
	}

	var deletedLaundry types.Laundry

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &deletedLaundry)
	deletedLaundry.DeletedAt = laundry.DeletedAt

	return deletedLaundry, nil
}

func (laundryRepository *LaundryRepository) Update(laundry *types.Laundry) (types.Laundry, error) {
	t := time.Now()
	laundry.UpdatedAt = &t

	id, _ := primitive.ObjectIDFromHex(laundry.ID)

	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{{"name", laundry.Name}, {"lat", laundry.Lat}, {"long", laundry.Long}}}}

	objectUpdated, err := laundryRepository.database.UpdateOne(collection, filter, update)
	if err != nil {
		panic(err)
	}

	if err != nil {
		return types.Laundry{}, err
	}

	var updatedLaundry types.Laundry

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &updatedLaundry)

	return types.Laundry{
		ID:        laundry.ID,
		Name:      laundry.Name,
		Lat:       laundry.Lat,
		Long:      laundry.Long,
		CreatedAt: updatedLaundry.CreatedAt,
		UpdatedAt: laundry.UpdatedAt,
	}, nil
}
