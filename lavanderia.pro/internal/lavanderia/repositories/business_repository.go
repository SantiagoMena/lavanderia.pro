package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/databases"
)

var businesscollection = "business"

type BusinessRepository struct {
	database databases.Database
}

func NewBusinessRepository(database databases.Database) *BusinessRepository {
	return &BusinessRepository{
		database: database,
	}
}

func (businessRepository *BusinessRepository) FindAllBusiness() ([]types.Business, error) {
	businessMap := []types.Business{}

	businessDb, err := businessRepository.database.FindAll(businesscollection)

	if err != nil {
		return nil, err
	}

	for businessDb.Next(context.TODO()) {
		var business types.Business

		if err := businessDb.Decode(&business); err != nil {
			return nil, err
		}

		businessMap = append(businessMap, business)
	}

	return businessMap, nil
}

func (businessRepository *BusinessRepository) Create(business *types.Business) (types.Business, error) {
	t := time.Now()
	business.CreatedAt = &t

	authId, _ := primitive.ObjectIDFromHex(business.Auth)

	businessDb, err := businessRepository.database.Create("business", bson.D{
		{Key: "name", Value: business.Name},
		{Key: "lat", Value: business.Lat},
		{Key: "long", Value: business.Long},
		{Key: "created_at", Value: business.CreatedAt},
		{Key: "auth", Value: authId},
	})

	if err != nil {
		return types.Business{}, err
	}

	insertedId := businessDb.InsertedID.(primitive.ObjectID).Hex()

	newBusiness := types.Business{
		ID:        insertedId,
		Name:      business.Name,
		Lat:       business.Lat,
		Long:      business.Long,
		CreatedAt: business.CreatedAt,
	}

	return newBusiness, nil
}

func (businessRepository *BusinessRepository) Delete(business *types.Business) (types.Business, error) {
	t := time.Now()
	business.DeletedAt = &t

	id, _ := primitive.ObjectIDFromHex(business.ID)

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"deleted_at", business.DeletedAt}}}}

	objectUpdated, err := businessRepository.database.UpdateOne(businesscollection, filter, update)
	if err != nil {
		panic(err)
	}

	if err != nil {
		return types.Business{}, err
	}

	var deletedBusiness types.Business

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &deletedBusiness)
	deletedBusiness.DeletedAt = business.DeletedAt

	return deletedBusiness, nil
}

func (businessRepository *BusinessRepository) Update(business *types.Business) (types.Business, error) {
	t := time.Now()
	business.UpdatedAt = &t

	id, _ := primitive.ObjectIDFromHex(business.ID)

	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{{"name", business.Name}, {"lat", business.Lat}, {"long", business.Long}}}}

	objectUpdated, err := businessRepository.database.UpdateOne(businesscollection, filter, update)
	if err != nil {
		panic(err)
	}

	if err != nil {
		return types.Business{}, err
	}

	var updatedBusiness types.Business

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &updatedBusiness)

	return types.Business{
		ID:        business.ID,
		Name:      business.Name,
		Lat:       business.Lat,
		Long:      business.Long,
		CreatedAt: updatedBusiness.CreatedAt,
		UpdatedAt: business.UpdatedAt,
	}, nil
}

func (businessRepository *BusinessRepository) Get(business *types.Business) (types.Business, error) {
	t := time.Now()
	business.UpdatedAt = &t

	id, _ := primitive.ObjectIDFromHex(business.ID)

	filter := bson.D{{"_id", id}}

	objectUpdated, err := businessRepository.database.FindOne(businesscollection, filter)
	if err != nil {
		return types.Business{}, err
	}

	if err != nil {
		return types.Business{}, err
	}

	var updatedBusiness types.Business

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &updatedBusiness)

	return types.Business{
		ID:        business.ID,
		Name:      business.Name,
		Lat:       business.Lat,
		Long:      business.Long,
		CreatedAt: updatedBusiness.CreatedAt,
		UpdatedAt: business.UpdatedAt,
		DeletedAt: business.DeletedAt,
	}, nil
}
