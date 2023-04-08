package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
	"lavanderia.pro/internal/lavanderia/handlers/client"
	"lavanderia.pro/internal/lavanderia/handlers/delivery"
)

type AuthController struct {
	RegisterBusinessHandler *business.RegisterBusinessHandler
	LoginHandler            *auth.LoginHandler
	RefreshTokenHandler     *auth.RefreshTokenHandler
	RegisterClientHandler   *client.RegisterClientHandler
	RegisterDeliveryHandler *delivery.RegisterDeliveryHandler
}

func NewAuthController(
	RegisterBusinessHandler *business.RegisterBusinessHandler,
	LoginHandler *auth.LoginHandler,
	RefreshTokenHandler *auth.RefreshTokenHandler,
	RegisterClientHandler *client.RegisterClientHandler,
	RegisterDeliveryHandler *delivery.RegisterDeliveryHandler,
) *AuthController {
	return &AuthController{
		RegisterBusinessHandler: RegisterBusinessHandler,
		LoginHandler:            LoginHandler,
		RefreshTokenHandler:     RefreshTokenHandler,
		RegisterClientHandler:   RegisterClientHandler,
		RegisterDeliveryHandler: RegisterDeliveryHandler,
	}
}

func (controller AuthController) RegisterBusiness(auth *types.Auth, business *types.Business) (types.Business, error) {
	// Handle Create Business
	businessDb, err := controller.RegisterBusinessHandler.Handle(auth, business)

	if err != nil {
		return types.Business{}, err
	}

	return businessDb, err
}

func (controller AuthController) Login(auth *types.Auth) (*types.JWT, error) {
	// Handle Create Business
	authDb, err := controller.LoginHandler.Handle(auth)

	if err != nil {
		return &types.JWT{}, err
	}

	return authDb, err
}

func (controller AuthController) RefreshToken(token string) (*types.JWT, error) {
	jwt, err := controller.RefreshTokenHandler.Handle(token)

	if err != nil {
		return &types.JWT{}, err
	}

	return jwt, err
}

func (controller AuthController) RegisterClient(auth *types.Auth, client *types.Client) (types.Client, error) {
	clientDb, err := controller.RegisterClientHandler.Handle(auth, client)

	if err != nil {
		return types.Client{}, err
	}

	return clientDb, err
}

func (controller AuthController) RegisterDelivery(auth *types.Auth, delivery *types.Delivery) (types.Delivery, error) {
	deliverytDb, err := controller.RegisterDeliveryHandler.Handle(auth, delivery)

	if err != nil {
		return types.Delivery{}, err
	}

	return deliverytDb, err
}
