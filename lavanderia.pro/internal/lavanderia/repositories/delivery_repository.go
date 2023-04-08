package repositories

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/databases"
)

var deliveryCollection = "delivery"

type DeliveryRepository struct {
	database databases.Database
}

func NewDeliveryRepository(database databases.Database) *DeliveryRepository {
	return &DeliveryRepository{
		database: database,
	}
}

func (deliveryRepository *DeliveryRepository) Create(delivery *types.Delivery) (types.Delivery, error) {
	t := time.Now()
	delivery.CreatedAt = &t

	authId, _ := primitive.ObjectIDFromHex(delivery.Auth)

	deliveryDb, err := deliveryRepository.database.Create("delivery", bson.D{
		{Key: "name", Value: delivery.Name},
		{Key: "created_at", Value: delivery.CreatedAt},
		{Key: "auth", Value: authId},
	})

	if err != nil {
		return types.Delivery{}, err
	}

	insertedId := deliveryDb.InsertedID.(primitive.ObjectID).Hex()

	newDelivery := types.Delivery{
		ID:        insertedId,
		Name:      delivery.Name,
		CreatedAt: delivery.CreatedAt,
	}

	return newDelivery, nil
}
