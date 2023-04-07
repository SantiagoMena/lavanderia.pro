package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/controllers"
)

func NewPostRegisterClientRouter(r *gin.Engine, controller *controllers.AuthController) {
	r.POST("/client/register", func(c *gin.Context) {

		var newClient types.Client
		var newAuth types.Auth

		// Call BindJSON to bind the received JSON to
		// newClient.
		if errClientJson := c.ShouldBindBodyWith(&newClient, binding.JSON); errClientJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errClientJson})
			return
		}

		// Call BindJSON to bind the received JSON to
		// newAuth.
		if errAuthJson := c.ShouldBindBodyWith(&newAuth, binding.JSON); errAuthJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errAuthJson})
			return
		}

		// Handle Controller
		client, errRegister := controller.RegisterClient(&newAuth, &newClient)

		if errRegister != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errRegister.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, client)
		}

	})
}

func NewGetClientRouter(r *gin.Engine, controller *controllers.ClientController) {
	r.GET("/client/profile", func(c *gin.Context) {
		authId := c.MustGet("auth")

		// Handle Controller
		client, errRegister := controller.GetClient(&types.Client{
			Auth: authId.(string),
		})

		if errRegister != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errRegister.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, client)
		}

	})
}
