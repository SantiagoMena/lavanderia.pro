package controllers

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/handlers/business"
)

type AuthBusinessController struct {
	RegisterBusinessHandler *business.RegisterBusinessHandler
}

func NewAuthBusinessController(
	RegisterBusinessHandler *business.RegisterBusinessHandler,
) *AuthBusinessController {
	return &AuthBusinessController{
		RegisterBusinessHandler: RegisterBusinessHandler,
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
