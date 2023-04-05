package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/controllers"
)

func NewPostLoginRouter(r *gin.Engine, controller *controllers.AuthController) {
	r.POST("/auth/login", func(c *gin.Context) {
		var newAuth types.Auth

		// Call BindJSON to bind the received JSON to
		// newAuth.
		if errAuthJson := c.ShouldBindBodyWith(&newAuth, binding.JSON); errAuthJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errAuthJson})
			return
		}

		// Handle Controller
		jwt, errRegister := controller.Login(&newAuth)

		if errRegister != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errRegister.Error()})
		} else {
			// TODO: return JWT
			c.IndentedJSON(http.StatusOK, jwt)
		}

	})
}

func NewPostRefreshTokenRouter(r *gin.Engine, controller *controllers.AuthController) {
	r.POST("/auth/refresh", func(c *gin.Context) {
		var token types.JWT

		// Call BindJSON to bind the received JSON to
		// newAuth.
		if errAuthJson := c.ShouldBindBodyWith(&token, binding.JSON); errAuthJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errAuthJson})
			return
		}

		// Handle Controller
		jwt, errRegister := controller.RefreshToken(token.RefreshToken)

		if errRegister != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errRegister.Error()})
		} else {
			// TODO: return JWT
			c.IndentedJSON(http.StatusOK, jwt)
		}

	})
}
