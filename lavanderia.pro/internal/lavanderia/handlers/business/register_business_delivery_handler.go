package business

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type RegisterBusinessDeliveryHandler struct {
	repositoryAuth     *repositories.AuthRepository
	repositoryBusiness *repositories.BusinessRepository
	repositoryDelivery *repositories.DeliveryRepository
}

func NewRegisterBusinessDeliveryHandler(
	repositoryAuth *repositories.AuthRepository,
	repositoryBusiness *repositories.BusinessRepository,
	repositoryDelivery *repositories.DeliveryRepository,
) *RegisterBusinessDeliveryHandler {
	return &RegisterBusinessDeliveryHandler{
		repositoryAuth:     repositoryAuth,
		repositoryBusiness: repositoryBusiness,
		repositoryDelivery: repositoryDelivery,
	}
}

func (ch RegisterBusinessDeliveryHandler) Handle(auth *types.Auth, business *types.Business, delivery *types.Delivery) (types.Delivery, error) {
	authFound, err := ch.repositoryAuth.GetByEmail(auth)

	if err != nil {
		return types.Delivery{}, errors.New("error on check auth")
	}

	if len(authFound.Email) > 0 {
		return types.Delivery{}, errors.New("auth already exists")
	}

	password, errPass := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)

	if errPass != nil {
		return types.Delivery{}, errors.New("error on encrypt password")
	}

	authDb, errCreateAuth := ch.repositoryAuth.Create(&types.Auth{
		Email:    auth.Email,
		Password: string(password),
	})

	if errCreateAuth != nil {
		return types.Delivery{}, errors.New("error on create auth")
	}

	deliveryDb, err := ch.repositoryDelivery.Create(&types.Delivery{
		Auth:     authDb.ID,
		Name:     delivery.Name,
		Business: business.ID,
	})

	return deliveryDb, err
}
