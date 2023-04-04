package business

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type RegisterBusinessHandler struct {
	repositoryAuth     *repositories.AuthRepository
	repositoryBusiness *repositories.BusinessRepository
}

func NewRegisterBusinessHandler(repositoryAuth *repositories.AuthRepository, repositoryBusiness *repositories.BusinessRepository) *RegisterBusinessHandler {
	return &RegisterBusinessHandler{
		repositoryAuth:     repositoryAuth,
		repositoryBusiness: repositoryBusiness,
	}
}

func (ch RegisterBusinessHandler) Handle(auth *types.Auth, business *types.Business) (types.Business, error) {
	authFound, err := ch.repositoryAuth.GetByEmail(auth)
	// authEmpty := types.Auth{}

	// panic(authFound)
	if len(authFound.Email) > 0 {
		return types.Business{}, errors.New("auth already exists")
	}

	password, errPass := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)

	if errPass != nil {
		return types.Business{}, errors.New("Error on encrypt password")
	}

	authDb, err := ch.repositoryAuth.Create(&types.Auth{
		Email:    auth.Email,
		Password: string(password),
	})

	businessDb, err := ch.repositoryBusiness.Create(&types.Business{
		Auth:      authDb.ID,
		Name:      business.Name,
		Lat:       business.Lat,
		Long:      business.Long,
		CreatedAt: business.CreatedAt,
	})

	return businessDb, err
}
