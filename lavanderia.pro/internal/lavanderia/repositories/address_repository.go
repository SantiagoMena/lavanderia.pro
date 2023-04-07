package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/databases"
)

var addressCollection = "address"

type AddressRepository struct {
	database databases.Database
}

func NewAddressRepository(database databases.Database) *AddressRepository {
	return &AddressRepository{
		database: database,
	}
}

func (addressRepository *AddressRepository) Create(address *types.Address) (types.Address, error) {
	t := time.Now()
	address.CreatedAt = &t

	clientId, _ := primitive.ObjectIDFromHex(address.Client)

	addressDb, err := addressRepository.database.Create("address", bson.D{
		{Key: "name", Value: address.Name},
		{Key: "position", Value: bson.D{
			{"type", "Point"},
			{"coordinates", address.Position},
		}},
		{Key: "address", Value: address.CreatedAt},
		{Key: "extra", Value: address.Extra},
		{Key: "client", Value: clientId},
		{Key: "created_at", Value: address.CreatedAt},
	})

	if err != nil {
		return types.Address{}, err
	}

	insertedId := addressDb.InsertedID.(primitive.ObjectID).Hex()

	newAddress := types.Address{
		ID:        insertedId,
		Name:      address.Name,
		Position:  address.Position,
		Address:   address.Address,
		Extra:     address.Extra,
		Phone:     address.Phone,
		Client:    address.Client,
		CreatedAt: address.CreatedAt,
	}

	return newAddress, nil
}

func (addressRepository *AddressRepository) Get(address *types.Address) (*types.Address, error) {
	id, _ := primitive.ObjectIDFromHex(address.ID)

	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "deleted_at", Value: nil},
	}

	objectAddress, err := addressRepository.database.FindOne(addressCollection, filter)

	if err != nil {
		return &types.Address{}, err
	}

	var foundAddress types.Address

	objectUpdt, _ := bson.Marshal(objectAddress)
	bson.Unmarshal(objectUpdt, &foundAddress)

	return &foundAddress, nil
}
