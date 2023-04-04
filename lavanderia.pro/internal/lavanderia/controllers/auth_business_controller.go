package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
)

type AuthBusinessController struct {
	RegisterBusinessHandler *business.RegisterBusinessHandler
	LoginHandler            *auth.LoginHandler
}

func NewAuthBusinessController(
	RegisterBusinessHandler *business.RegisterBusinessHandler,
	LoginHandler *auth.LoginHandler,
) *AuthBusinessController {
	return &AuthBusinessController{
		RegisterBusinessHandler: RegisterBusinessHandler,
		LoginHandler:            LoginHandler,
	}
}

func (controller AuthBusinessController) RegisterBusiness(auth *types.Auth, business *types.Business) (types.Business, error) {
	// Handle Create Business
	businessDb, err := controller.RegisterBusinessHandler.Handle(auth, business)

	if err != nil {
		return types.Business{}, err
	}

	return businessDb, err
}

func (controller AuthBusinessController) Login(auth *types.Auth) (*types.JWT, error) {
	// Handle Create Business
	authDb, err := controller.LoginHandler.Handle(auth)

	if err != nil {
		return &types.JWT{}, err
	}

	return authDb, err
}
