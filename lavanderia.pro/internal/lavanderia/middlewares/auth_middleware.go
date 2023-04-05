package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"lavanderia.pro/internal/lavanderia/repositories"
)

type Header struct {
	Token string `header:"Authorization" binding:"required"`
}

func NewJWTMiddleware(r *gin.Engine, authRepository *repositories.AuthRepository) {
	r.Use(SetAuthJWT(authRepository))
}

func SetAuthJWT(authRepository *repositories.AuthRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		header := &Header{}

		// bind the headers to data
		if err := c.ShouldBindHeader(header); err != nil {
			// c.JSON(http.StatusForbidden, err.Error())
			// return
		}

		token := strings.Split(header.Token, " ")

		if len(token) > 1 {
			auth, errAuth := authRepository.GetAuthByToken(token[1])
			if errAuth != nil {
				c.JSON(http.StatusForbidden, gin.H{"msg": errAuth.Error()})
				return
			}
			c.Set("auth", auth.ID)

		} else {
			c.Set("auth", nil)
		}

		c.Next()
	}
}
