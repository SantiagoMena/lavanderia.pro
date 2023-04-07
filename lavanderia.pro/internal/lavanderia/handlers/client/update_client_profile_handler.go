package client

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type UpdateClientProfileHandler struct {
	repository *repositories.ClientRepository
}

func NewUpdateClientProfileHandler(clientRepository *repositories.ClientRepository) *UpdateClientProfileHandler {
	return &UpdateClientProfileHandler{
		repository: clientRepository,
	}
}

func (ch *UpdateClientProfileHandler) Handle(client *types.Client) (types.Client, error) {
	// find client
	clientFound, errorFind := ch.repository.GetClientByAuth(client)
	if errorFind != nil {
		return types.Client{}, errorFind
	}

	clientFound.Name = client.Name

	// update client
	clientUpdated, errorUpdate := ch.repository.Update(&clientFound)

	return clientUpdated, errorUpdate
}
