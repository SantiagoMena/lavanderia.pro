package client

import (
	"errors"

	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type PostClientHandler struct {
	repository *repositories.ClientRepository
}

func NewPostClientHandler(clientRepository *repositories.ClientRepository) *PostClientHandler {
	return &PostClientHandler{
		repository: clientRepository,
	}
}

func (ch *PostClientHandler) Handle(client *types.Client) (*types.Client, error) {
	// find client by auth
	clientFound, _ := ch.repository.GetClientByAuth(client)
	emptyClient := types.Client{}

	if clientFound != emptyClient {
		return &types.Client{}, errors.New("client already registered")
	}

	// if exists error

	// if not create

	// return client

	return &types.Client{}, nil
}
