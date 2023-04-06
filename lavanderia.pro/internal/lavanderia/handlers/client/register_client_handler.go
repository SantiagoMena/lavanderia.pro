package client

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type RegisterClientsHandler struct {
	repositoryAuth   *repositories.AuthRepository
	repositoryClient *repositories.ClientRepository
}

func NewRegisterClientsHandler(repositoryAuth *repositories.AuthRepository, repositoryClient *repositories.ClientRepository) *RegisterClientsHandler {
	return &RegisterClientsHandler{
		repositoryAuth:   repositoryAuth,
		repositoryClient: repositoryClient,
	}
}

func (ch RegisterClientsHandler) Handle(auth *types.Auth, business *types.Client) (types.Client, error) {
	authFound, err := ch.repositoryAuth.GetByEmail(auth)

	// panic(authFound)
	if len(authFound.Email) > 0 {
		return types.Client{}, errors.New("auth already exists")
	}

	password, errPass := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)

	if errPass != nil {
		return types.Client{}, errors.New("Error on encrypt password")
	}

	authDb, err := ch.repositoryAuth.Create(&types.Auth{
		Email:    auth.Email,
		Password: string(password),
	})

	clientDb, err := ch.repositoryClient.Create(&types.Client{
		Auth:      authDb.ID,
		Name:      business.Name,
		CreatedAt: business.CreatedAt,
	})

	return clientDb, err
}
