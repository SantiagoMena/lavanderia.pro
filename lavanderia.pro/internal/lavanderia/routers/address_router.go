package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/controllers"
	"lavanderia.pro/internal/lavanderia/repositories"
)

func NewPostAddressRouter(r *gin.Engine, controller *controllers.AddressController, clientRepository *repositories.ClientRepository) {
	r.POST("/address", func(c *gin.Context) {
		authId := c.MustGet("auth")

		var newAddress types.Address
		if err := c.BindJSON(&newAddress); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		client, errClient := clientRepository.GetClientByAuth(&types.Client{
			Auth: authId.(string),
		})

		if errClient != nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": errClient})
			return
		}

		address, errAddress := controller.CreateAddress(&types.Address{
			Client:   client.ID,
			Name:     newAddress.Name,
			Position: newAddress.Position,
			Address:  newAddress.Address,
			Extra:    newAddress.Extra,
		})

		if errAddress != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errAddress.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, address)
		}
	})
}
