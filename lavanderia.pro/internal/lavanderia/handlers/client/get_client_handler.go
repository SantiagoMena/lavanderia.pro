package client

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type GetClientHandler struct {
	repositoryAuth   *repositories.AuthRepository
	repositoryClient *repositories.ClientRepository
}

func NewGetClientHandler(repositoryAuth *repositories.AuthRepository, repositoryClient *repositories.ClientRepository) *GetClientHandler {
	return &GetClientHandler{
		repositoryAuth:   repositoryAuth,
		repositoryClient: repositoryClient,
	}
}

func (ch GetClientHandler) Handle(client *types.Client) (types.Client, error) {
	clientDb, err := ch.repositoryClient.GetClient(client)

	return clientDb, err
}
