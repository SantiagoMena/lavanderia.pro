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

func (clientRepository *ClientRepository) Create(business *types.Client) (types.Client, error) {
	t := time.Now()
	business.CreatedAt = &t

	authId, _ := primitive.ObjectIDFromHex(business.Auth)

	clientDb, err := clientRepository.database.Create(clientCollection, bson.D{
		{Key: "name", Value: business.Name},
		{Key: "auth", Value: authId},
		{Key: "created_at", Value: business.CreatedAt},
	})

	if err != nil {
		return types.Client{}, err
	}

	insertedId := clientDb.InsertedID.(primitive.ObjectID).Hex()

	newClient := types.Client{
		ID:        insertedId,
		Name:      business.Name,
		CreatedAt: business.CreatedAt,
	}

	return newClient, nil
}

func (clientRepository *ClientRepository) GetClient(client *types.Client) (types.Client, error) {
	id, _ := primitive.ObjectIDFromHex(client.ID)

	clientFound, errFind := clientRepository.database.FindOne(clientCollection, bson.D{
		{Key: "_id", Value: id},
		{Key: "deleted_at", Value: nil},
	})

	if errFind != nil {
		return types.Client{}, errFind
	}

	var clientUnmarshal types.Client
	marshalObject, errMarshal := bson.Marshal(clientFound)
	bson.Unmarshal(marshalObject, &clientUnmarshal)

	if errMarshal != nil {
		return types.Client{}, errFind
	}

	return clientUnmarshal, nil
}
