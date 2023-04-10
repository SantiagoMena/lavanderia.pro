package delivery

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type RegisterDeliveryHandler struct {
	repositoryAuth     *repositories.AuthRepository
	repositoryDelivery *repositories.DeliveryRepository
}

func NewRegisterDeliveryHandler(repositoryAuth *repositories.AuthRepository, repositoryDelivery *repositories.DeliveryRepository) *RegisterDeliveryHandler {
	return &RegisterDeliveryHandler{
		repositoryAuth:     repositoryAuth,
		repositoryDelivery: repositoryDelivery,
	}
}

func (ch RegisterDeliveryHandler) Handle(auth *types.Auth, delivery *types.Delivery) (types.Delivery, error) {
	authFound, err := ch.repositoryAuth.GetByEmail(auth)
	// authEmpty := types.Auth{}

	// panic(authFound)
	if len(authFound.Email) > 0 {
		return types.Delivery{}, errors.New("auth already exists")
	}

	password, errPass := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)

	if errPass != nil {
		return types.Delivery{}, errors.New("Error on encrypt password")
	}

	authDb, err := ch.repositoryAuth.Create(&types.Auth{
		Email:    auth.Email,
		Password: string(password),
	})

	if err != nil {
		return types.Delivery{}, err
	}

	deliveryDb, err := ch.repositoryDelivery.Create(&types.Delivery{
		Auth:      authDb.ID,
		Name:      delivery.Name,
		Business:  delivery.Business,
		CreatedAt: delivery.CreatedAt,
	})

	return deliveryDb, err
}
