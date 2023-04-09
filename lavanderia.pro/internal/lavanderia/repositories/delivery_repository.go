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

func (deliveryRepository *DeliveryRepository) GetDeliveryByAuth(delivery *types.Delivery) (types.Delivery, error) {
	authId, _ := primitive.ObjectIDFromHex(delivery.Auth)

	deliveryFound, errFind := deliveryRepository.database.FindOne(deliveryCollection, bson.D{
		{Key: "auth", Value: authId},
		{Key: "deleted_at", Value: nil},
	})

	if errFind != nil {
		return types.Delivery{}, errFind
	}

	var deliveryUnmarshal types.Delivery
	marshalObject, errMarshal := bson.Marshal(deliveryFound)
	bson.Unmarshal(marshalObject, &deliveryUnmarshal)

	if errMarshal != nil {
		return types.Delivery{}, errFind
	}

	return deliveryUnmarshal, nil
}

func (deliveryRepository *DeliveryRepository) Update(delivery *types.Delivery) (types.Delivery, error) {
	t := time.Now()
	delivery.UpdatedAt = &t

	id, _ := primitive.ObjectIDFromHex(delivery.ID)

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: delivery.Name},
		}},
	}

	objectUpdated, err := deliveryRepository.database.UpdateOne(deliveryCollection, filter, update)
	if err != nil {
		return types.Delivery{}, err
	}

	var updatedDelivery types.Delivery

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &updatedDelivery)

	return types.Delivery{
		ID:        delivery.ID,
		Name:      delivery.Name,
		CreatedAt: updatedDelivery.CreatedAt,
		UpdatedAt: delivery.UpdatedAt,
	}, nil
}
