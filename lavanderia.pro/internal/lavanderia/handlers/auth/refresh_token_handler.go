package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type RefreshTokenHandler struct {
	repositoryAuth     *repositories.AuthRepository
	repositoryBusiness *repositories.BusinessRepository
}

func NewRefreshTokenHandler(
	repositoryAuth *repositories.AuthRepository,
	repositoryBusiness *repositories.BusinessRepository,
) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		repositoryAuth:     repositoryAuth,
		repositoryBusiness: repositoryBusiness,
	}
}

func (ch RefreshTokenHandler) Handle(auth *types.JWT) (*types.JWT, error) {
	token, err := jwt.Parse(auth.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_JWT")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["id"])

		return &types.JWT{
			Token: token.Raw,
		}, nil
	} else {
		return &types.JWT{}, err
	}
}
