package auth

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type CreateJWTHandler struct {
	repository *repositories.AuthRepository
}

func NewCreateJWTHandler(repository *repositories.AuthRepository) *CreateJWTHandler {
	return &CreateJWTHandler{repository: repository}
}

func (ch CreateJWTHandler) Handle(auth *types.Auth) (types.JWT, error) {
	jwt, err := ch.repository.CreateJWT(auth)

	return jwt, err
}
