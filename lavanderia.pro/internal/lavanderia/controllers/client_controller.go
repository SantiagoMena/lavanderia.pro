package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/client"
)

type ClientController struct {
	RegisterClientHandler *client.RegisterClientHandler
}

func NewClientController(
	RegisterClientHandler *client.RegisterClientHandler,
) *ClientController {
	return &ClientController{
		RegisterClientHandler: RegisterClientHandler,
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
