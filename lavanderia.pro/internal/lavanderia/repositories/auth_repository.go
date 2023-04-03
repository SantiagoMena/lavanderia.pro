package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/databases"
)

var authCollection = "auth"

type AuthRepository struct {
	database databases.Database
}

func NewAuthRepository(database databases.Database) *AuthRepository {
	return &AuthRepository{
		database: database,
	}
}

func (authRepository *AuthRepository) Create(auth *types.Auth) (types.Auth, error) {
	t := time.Now()
	auth.CreatedAt = &t

	authDb, err := authRepository.database.Create(authCollection, bson.D{
		{Key: "email", Value: auth.Email},
		{Key: "password", Value: auth.Password},
		{Key: "facebook_id", Value: auth.FacebookId},
		{Key: "google_id", Value: auth.GoogleId},
		{Key: "apple_id", Value: auth.AppleId},
		{Key: "created_at", Value: auth.CreatedAt},
	})

	if err != nil {
		return types.Auth{}, err
	}

	insertedId := authDb.InsertedID.(primitive.ObjectID).Hex()

	newAuth := types.Auth{
		ID:         insertedId,
		Email:      auth.Email,
		Password:   auth.Password,
		FacebookId: auth.FacebookId,
		GoogleId:   auth.GoogleId,
		AppleId:    auth.AppleId,
		CreatedAt:  auth.CreatedAt,
	}

	return newAuth, nil
}

func (authRepository *AuthRepository) GetByEmail(auth *types.Auth) (types.Auth, error) {
	filter := bson.D{{"email", auth.Email}}
	var emptyAuth types.Auth

	object, err := authRepository.database.FindOne(authCollection, filter)
	if err != nil {
		return emptyAuth, err
	}

	var foundAuth types.Auth

	objectAuth, _ := bson.Marshal(object)
	bson.Unmarshal(objectAuth, &foundAuth)

	return types.Auth{
		ID:    auth.ID,
		Email: auth.Email,
		// Password:   auth.Password,
		FacebookId: auth.FacebookId,
		GoogleId:   auth.GoogleId,
		AppleId:    auth.AppleId,
		CreatedAt:  auth.CreatedAt,
	}, nil
}
