package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/client"
)

type ClientController struct {
	RegisterClientHandler *client.RegisterClientHandler
	GetClientHandler      *client.GetClientHandler
}

func NewClientController(
	RegisterClientHandler *client.RegisterClientHandler,
	GetClientHandler *client.GetClientHandler,
) *ClientController {
	return &ClientController{
		RegisterClientHandler: RegisterClientHandler,
		GetClientHandler:      GetClientHandler,
	}
}

func (controller ClientController) RegisterClient(auth *types.Auth, client *types.Client) (types.Client, error) {
	// Handle Create Client
	clientDb, err := controller.RegisterClientHandler.Handle(auth, client)

	if err != nil {
		return types.Client{}, err
	}

	return clientDb, err
}

func (controller ClientController) GetClient(client *types.Client) (types.Client, error) {
	// Handle Create Client
	clientDb, err := controller.GetClientHandler.Handle(client)

	if err != nil {
		return types.Client{}, err
	}

	return clientDb, err
}
