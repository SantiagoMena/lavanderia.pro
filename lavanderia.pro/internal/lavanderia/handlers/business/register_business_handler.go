package business

import (
	"errors"

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
	authEmpty := types.Auth{}

	if authFound != authEmpty {
		return types.Business{}, errors.New("auth already exists")
	}

	authDb, err := ch.repositoryAuth.Create(auth)

	businessDb, err := ch.repositoryBusiness.Create(&types.Business{
		Auth:      authDb.ID,
		Name:      business.Name,
		Lat:       business.Lat,
		Long:      business.Long,
		CreatedAt: business.CreatedAt,
	})

	return businessDb, err
}
