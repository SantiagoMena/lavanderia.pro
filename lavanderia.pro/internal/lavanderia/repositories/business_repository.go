package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/databases"
)

var businessCollection = "business"

type BusinessRepository struct {
	database databases.Database
}

func NewBusinessRepository(database databases.Database) *BusinessRepository {
	return &BusinessRepository{
		database: database,
	}
}

func (businessRepository *BusinessRepository) FindAllBusiness() ([]types.Business, error) {
	filter := bson.D{
		{Key: "deleted_at", Value: nil},
	}

	businessDb, err := businessRepository.database.FindAllFilter(businessCollection, filter)

	if err != nil {
		return nil, err
	}

	var businessMap []types.Business
	if err = businessDb.All(context.TODO(), &businessMap); err != nil {
		panic(err)
	}

	return businessMap, nil
}

func (businessRepository *BusinessRepository) Create(business *types.Business) (types.Business, error) {
	t := time.Now()
	business.CreatedAt = &t

	authId, _ := primitive.ObjectIDFromHex(business.Auth)

	businessDb, err := businessRepository.database.Create("business", bson.D{
		{Key: "name", Value: business.Name},
		{Key: "position", Value: business.Position},
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
		Position:  business.Position,
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

	objectUpdated, err := businessRepository.database.UpdateOne(businessCollection, filter, update)
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
		{"$set", bson.D{
			{"name", business.Name},
			{"position", business.Position},
		}}}

	objectUpdated, err := businessRepository.database.UpdateOne(businessCollection, filter, update)
	if err != nil {
		return types.Business{}, err
	}

	var updatedBusiness types.Business

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &updatedBusiness)

	return types.Business{
		ID:        business.ID,
		Name:      business.Name,
		Position:  business.Position,
		CreatedAt: updatedBusiness.CreatedAt,
		UpdatedAt: business.UpdatedAt,
	}, nil
}

func (businessRepository *BusinessRepository) Get(business *types.Business) (types.Business, error) {
	t := time.Now()
	business.UpdatedAt = &t

	id, _ := primitive.ObjectIDFromHex(business.ID)

	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "deleted_at", Value: nil},
	}

	objectBusiness, err := businessRepository.database.FindOne(businessCollection, filter)
	if err != nil {
		return types.Business{}, err
	}

	if err != nil {
		return types.Business{}, err
	}

	var foundBusiness types.Business

	objectUpdt, _ := bson.Marshal(objectBusiness)
	bson.Unmarshal(objectUpdt, &foundBusiness)

	return foundBusiness, nil
}

func (businessRepository *BusinessRepository) FindAllBusinessByAuth(auth string) ([]types.Business, error) {
	// businessMap := []types.Business{}

	authID, _ := primitive.ObjectIDFromHex(auth)

	businessDb, err := businessRepository.database.FindAllFilter(businessCollection, bson.D{
		{Key: "auth", Value: authID},
	})

	if err != nil {
		return nil, err
	}

	var businessMap []types.Business
	if err = businessDb.All(context.TODO(), &businessMap); err != nil {
		return []types.Business{}, err
	}

	return businessMap, nil
}
