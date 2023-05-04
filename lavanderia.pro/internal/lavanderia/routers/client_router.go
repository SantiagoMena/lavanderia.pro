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

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		// Handle Controller
		client, errRegister := controller.GetClientByAuth(&types.Client{
			Auth: authId.(string),
		})

		if errRegister != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errRegister.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, client)
		}

	})
}

func NewPostClientRouter(r *gin.Engine, controller *controllers.ClientController) {
	r.POST("/client/profile", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		var client types.Client

		if err := c.BindJSON(&client); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		client.Auth = authId.(string)

		// Handle Controller
		clientPosted, errPost := controller.PostClient(&client)

		if errPost != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errPost.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, clientPosted)
		}

	})
}

func NewPutClientRouter(r *gin.Engine, controller *controllers.ClientController) {
	r.PUT("/client/profile", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		var client types.Client

		if err := c.BindJSON(&client); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		client.Auth = authId.(string)

		// Handle Controller
		clientPosted, errPost := controller.UpdateClient(&client)

		if errPost != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errPost.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, clientPosted)
		}

	})
}
