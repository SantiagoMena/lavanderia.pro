package repositories

import (
	"errors"

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
		{Key: "position", Value: address.Position},
		{Key: "address", Value: address.Address},
		{Key: "extra", Value: address.Extra},
		{Key: "client", Value: clientId},
		{Key: "phone", Value: address.Phone},
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

func (addressRepository *AddressRepository) Update(address *types.Address) (*types.Address, error) {
	id, _ := primitive.ObjectIDFromHex(address.ID)
	t := time.Now()
	address.UpdatedAt = &t

	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "deleted_at", Value: nil},
	}

	objectAddress, errFind := addressRepository.database.FindOne(addressCollection, filter)

	var foundAddress types.Address
	objectUpdt, _ := bson.Marshal(objectAddress)
	bson.Unmarshal(objectUpdt, &foundAddress)
	if errFind != nil {
		return &types.Address{}, errors.New("object not found")
	}

	update := bson.D{{"$set", bson.D{
		{Key: "client", Value: address.Client},
		{Key: "name", Value: address.Name},
		{Key: "position", Value: address.Position},
		{Key: "address", Value: address.Address},
		{Key: "phone", Value: address.Phone},
		{Key: "extra", Value: address.Extra},
		{Key: "updated_at", Value: address.UpdatedAt},
	}}}

	addressUpdatedMongo, errUpdate := addressRepository.database.UpdateOne(addressCollection, filter, update)

	if errUpdate != nil {
		return &types.Address{}, errors.New("error on update object")
	}

	var addressUpdated types.Address
	addressMarshal, _ := bson.Marshal(addressUpdatedMongo)
	bson.Unmarshal(addressMarshal, &addressUpdated)

	adressFound, errFind := addressRepository.database.UpdateOne(addressCollection, filter, update)

	var addressFoundUnmarshal types.Address
	addressFoundMarshal, _ := bson.Marshal(adressFound)
	bson.Unmarshal(addressFoundMarshal, &addressFoundUnmarshal)

	return &addressFoundUnmarshal, errFind
}
