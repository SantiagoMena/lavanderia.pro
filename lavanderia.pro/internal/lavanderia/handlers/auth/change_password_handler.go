package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type ChangePasswordHandler struct {
	repositoryAuth     *repositories.AuthRepository
	repositoryBusiness *repositories.BusinessRepository
}

func NewChangePasswordHandler(
	repositoryAuth *repositories.AuthRepository,
	repositoryBusiness *repositories.BusinessRepository,
) *ChangePasswordHandler {
	return &ChangePasswordHandler{
		repositoryAuth:     repositoryAuth,
		repositoryBusiness: repositoryBusiness,
	}
}

func (ch ChangePasswordHandler) Handle(authId string, newPassword *types.NewPassword) (types.Auth, error) {
	auth, errAuth := ch.repositoryAuth.GetById(authId)
	if errAuth != nil {
		return types.Auth{}, errAuth
	}

	hash := auth.Password
	password := newPassword.Password

	pass, er := validateCredentialsChangePassword([]byte(hash), []byte(password))

	if er != nil {
		return types.Auth{}, er
	}

	if pass != true {
		return types.Auth{}, errors.New("password incorrect")
	}

	// Register Client
	pwd := []byte(newPassword.NewPassword)
	passwordClient, errPassClient := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if errPassClient != nil {
		return types.Auth{}, errors.New("error encrypt password")
	}

	auth.Password = string(passwordClient)
	authUpdated, errAuthUpdate := ch.repositoryAuth.UpdatePassword(&auth)

	if errAuthUpdate != nil {
		return types.Auth{}, errors.New("error update password")
	}

	return authUpdated, nil
}

func validateCredentialsChangePassword(actualPasswordHash []byte, attemptedPassword []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(actualPasswordHash, attemptedPassword); err != nil {
		return false, nil
	}

	return true, nil
}
