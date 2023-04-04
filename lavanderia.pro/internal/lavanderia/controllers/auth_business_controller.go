package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/auth"
	"lavanderia.pro/internal/lavanderia/handlers/business"
)

type AuthBusinessController struct {
	RegisterBusinessHandler *business.RegisterBusinessHandler
	LoginBusinessHandler    *auth.LoginBusinessHandler
}

func NewAuthBusinessController(
	RegisterBusinessHandler *business.RegisterBusinessHandler,
	LoginBusinessHandler *auth.LoginBusinessHandler,
) *AuthBusinessController {
	return &AuthBusinessController{
		RegisterBusinessHandler: RegisterBusinessHandler,
		LoginBusinessHandler:    LoginBusinessHandler,
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

func (controller AuthBusinessController) LoginBusiness(auth *types.Auth) (*types.JWT, error) {
	// Handle Create Business
	authDb, err := controller.LoginBusinessHandler.Handle(auth)

	if err != nil {
		return &types.JWT{}, err
	}

	return authDb, err
}
