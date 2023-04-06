package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/databases"
)

var clientCollection = "client"

type ClientRepository struct {
	database databases.Database
}

func NewClientRepository(database databases.Database) *ClientRepository {
	return &ClientRepository{
		database: database,
	}
}

func (businessRepository *ClientRepository) Create(business *types.Client) (types.Client, error) {
	t := time.Now()
	business.CreatedAt = &t

	authId, _ := primitive.ObjectIDFromHex(business.Auth)

	businessDb, err := businessRepository.database.Create("business", bson.D{
		{Key: "name", Value: business.Name},
		{Key: "auth", Value: authId},
		{Key: "created_at", Value: business.CreatedAt},
	})

	if err != nil {
		return types.Client{}, err
	}

	insertedId := businessDb.InsertedID.(primitive.ObjectID).Hex()

	newClient := types.Client{
		ID:        insertedId,
		Name:      business.Name,
		CreatedAt: business.CreatedAt,
	}

	return newClient, nil
}
