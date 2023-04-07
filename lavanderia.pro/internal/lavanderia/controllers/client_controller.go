package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/client"
)

type ClientController struct {
	RegisterClientHandler *client.RegisterClientHandler
	GetClientHandler      *client.GetClientHandler
	PostClientHandler     *client.PostClientHandler
}

func NewClientController(
	RegisterClientHandler *client.RegisterClientHandler,
	GetClientHandler *client.GetClientHandler,
	PostClientHandler *client.PostClientHandler,
) *ClientController {
	return &ClientController{
		RegisterClientHandler: RegisterClientHandler,
		GetClientHandler:      GetClientHandler,
		PostClientHandler:     PostClientHandler,
	}
}

func (controller ClientController) RegisterClient(auth *types.Auth, client *types.Client) (types.Client, error) {
	clientDb, err := controller.RegisterClientHandler.Handle(auth, client)

	if err != nil {
		return types.Client{}, err
	}

	return clientDb, err
}

func (controller ClientController) GetClientByAuth(client *types.Client) (types.Client, error) {
	clientDb, err := controller.GetClientHandler.Handle(client)

	if err != nil {
		return types.Client{}, err
	}

	return clientDb, err
}

func (controller ClientController) PostClient(client *types.Client) (*types.Client, error) {
	clientDb, err := controller.PostClientHandler.Handle(client)

	if err != nil {
		return &types.Client{}, err
	}

	return clientDb, err
}
