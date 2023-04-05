package auth

import (
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type RefreshTokenHandler struct {
	repositoryAuth *repositories.AuthRepository
}

func NewRefreshTokenHandler(
	repositoryAuth *repositories.AuthRepository,
) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		repositoryAuth: repositoryAuth,
	}
}

func (ch RefreshTokenHandler) Handle(refreshToken string) (*types.JWT, error) {
	token, err := ch.repositoryAuth.RefreshJWT(refreshToken)

	return token, err
}
