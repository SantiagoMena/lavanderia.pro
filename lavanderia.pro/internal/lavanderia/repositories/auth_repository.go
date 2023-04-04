package repositories

import (
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
	mySigningKey := []byte("SECRET_JWT_SIGN_KEY")

	type CustomClaims struct {
		Auth *types.Auth `json:"auth"`
		jwt.RegisteredClaims
	}

	tokenExpires := time.Now().Add(24 * time.Hour)
	// Create claims while leaving out some of the optional fields
	claims := CustomClaims{
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

	type CustomClaimsRefresh struct {
		jwt.RegisteredClaims
	}

	refreshTokenExpires := time.Now().Add(24 * time.Hour * 30)
	claimsRefresh := CustomClaimsRefresh{
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
