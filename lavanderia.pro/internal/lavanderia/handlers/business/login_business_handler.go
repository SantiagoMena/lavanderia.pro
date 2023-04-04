package business

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type LoginBusinessHandler struct {
	repositoryAuth     *repositories.AuthRepository
	repositoryBusiness *repositories.BusinessRepository
}

func NewLoginBusinessHandler(
	repositoryAuth *repositories.AuthRepository,
	repositoryBusiness *repositories.BusinessRepository,
) *LoginBusinessHandler {
	return &LoginBusinessHandler{
		repositoryAuth:     repositoryAuth,
		repositoryBusiness: repositoryBusiness,
	}
}

func (ch LoginBusinessHandler) Handle(auth *types.Auth) (types.Auth, error) {
	authFound, err := ch.repositoryAuth.GetByEmail(&types.Auth{
		Email: auth.Email,
	})
	if err != nil {
		return types.Auth{}, errors.New("email incorrect")
	}

	hash := authFound.Password
	password := auth.Password

	pass, er := validateCredentials([]byte(hash), []byte(password))

	if er != nil {
		return types.Auth{}, er
	}

	if pass != true {
		return types.Auth{}, errors.New("password incorrect")
	}

	return authFound, nil
}

func validateCredentials(actualPasswordHash []byte, attemptedPassword []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(actualPasswordHash, attemptedPassword); err != nil {
		return false, nil
	}

	return true, nil
}
