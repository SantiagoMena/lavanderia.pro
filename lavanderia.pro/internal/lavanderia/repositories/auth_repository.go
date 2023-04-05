package repositories

import (
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (authRepository *AuthRepository) GetById(id string) (types.Auth, error) {
	var emptyAuth types.Auth
	ObjectID, errOBjIdFromHex := primitive.ObjectIDFromHex(id)

	if errOBjIdFromHex != nil {
		return emptyAuth, errOBjIdFromHex
	}
	filter := bson.D{{"_id", ObjectID}}

	object, err := authRepository.database.FindOne(authCollection, filter)
	if err != nil {
		return emptyAuth, err
	}

	var foundAuth types.Auth

	objectAuth, _ := bson.Marshal(object)
	bson.Unmarshal(objectAuth, &foundAuth)

	return types.Auth{
		ID:         foundAuth.ID,
		Email:      foundAuth.Email,
		Password:   foundAuth.Password,
		FacebookId: foundAuth.FacebookId,
		GoogleId:   foundAuth.GoogleId,
		AppleId:    foundAuth.AppleId,
		CreatedAt:  foundAuth.CreatedAt,
	}, nil
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
		ID:         foundAuth.ID,
		Email:      foundAuth.Email,
		Password:   foundAuth.Password,
		FacebookId: foundAuth.FacebookId,
		GoogleId:   foundAuth.GoogleId,
		AppleId:    foundAuth.AppleId,
		CreatedAt:  foundAuth.CreatedAt,
	}, nil
}

func (authRepository *AuthRepository) CreateJWT(auth *types.Auth) (*types.JWT, error) {
	auth = &types.Auth{
		ID:         auth.ID,
		Email:      auth.Email,
		GoogleId:   auth.GoogleId,
		FacebookId: auth.FacebookId,
		AppleId:    auth.AppleId,
	}

	// mySigningKey := []byte(auth.Password)
	mySigningKey := []byte(os.Getenv("SECRET_JWT"))

	type CustomClaims struct {
		Type string      `json:"type"`
		Auth *types.Auth `json:"auth"`
		jwt.RegisteredClaims
	}

	tokenExpires := time.Now().Add(24 * time.Hour)
	// Create claims while leaving out some of the optional fields
	claims := CustomClaims{
		"token",
		auth,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(tokenExpires),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, errSignToken := token.SignedString(mySigningKey)

	if errSignToken != nil {
		return &types.JWT{}, errSignToken
	}

	refreshTokenExpires := time.Now().Add(24 * time.Hour * 30)
	claimsRefresh := CustomClaims{
		"refresh_token",
		auth,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpires),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	refreshTokenSigned, errSignRefreshToken := refreshToken.SignedString(mySigningKey)

	if errSignRefreshToken != nil {
		return &types.JWT{}, errSignRefreshToken
	}

	return &types.JWT{
		Token:        tokenSigned,
		RefreshToken: refreshTokenSigned,
	}, nil
}

func (authRepository *AuthRepository) RefreshJWT(refreshToken string) (*types.JWT, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_JWT")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenType, _ := claims["type"].(string)
		if tokenType != "refresh_token" {
			return &types.JWT{}, errors.New("token type error")
		}

		auth, _ := claims["auth"]
		authMap, _ := auth.(map[string]interface{})

		authObj, errGetAuth := authRepository.GetById(authMap["id"].(string))

		if errGetAuth != nil {
			return nil, errGetAuth
		}

		refreshedToken, errRefresh := authRepository.CreateJWT(&authObj)

		return refreshedToken, errRefresh
	} else {
		return &types.JWT{}, err
	}
}
