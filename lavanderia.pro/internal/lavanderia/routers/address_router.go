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

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
		}

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
			Phone:    newAddress.Phone,
			Extra:    newAddress.Extra,
		})

		if errAddress != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errAddress.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, address)
		}
	})
}

func NewGetAddressRouter(r *gin.Engine, controller *controllers.AddressController, clientRepository *repositories.ClientRepository) {
	r.GET("/address/:id", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
		}

		var addressId types.Address
		if err := c.ShouldBindUri(&addressId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		client, errClient := clientRepository.GetClientByAuth(&types.Client{
			Auth: authId.(string),
		})

		if errClient != nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": errClient.Error()})
			return
		}

		address, errAddress := controller.GetAddress(&addressId)

		if errAddress != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"msg": errAddress.Error()})
			return
		}

		if address.Client != client.ID {
			c.IndentedJSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		c.IndentedJSON(http.StatusCreated, address)
	})
}

func NewUpdateAddressRouter(r *gin.Engine, controller *controllers.AddressController, clientRepository *repositories.ClientRepository) {
	r.PUT("/address/:id", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
		}
		var addressId types.Address

		if err := c.ShouldBindUri(&addressId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

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

		address, errAddress := controller.UpdateAddress(&types.Address{
			ID:       addressId.ID,
			Client:   client.ID,
			Name:     newAddress.Name,
			Position: newAddress.Position,
			Address:  newAddress.Address,
			Phone:    newAddress.Phone,
			Extra:    newAddress.Extra,
		})

		if errAddress != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errAddress.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, address)
		}
	})
}

func NewGetAddressesRouter(r *gin.Engine, controller *controllers.AddressController, clientRepository *repositories.ClientRepository) {
	r.GET("/addresses", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
		}

		client, errClient := clientRepository.GetClientByAuth(&types.Client{
			Auth: authId.(string),
		})

		if errClient != nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": errClient.Error()})
			return
		}

		addresses, errAddress := controller.GetAddresses(&types.Address{
			Client: client.ID,
		})

		if errAddress != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"msg": errAddress.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, addresses)
	})
}

func NewDeleteAddressRouter(r *gin.Engine, controller *controllers.AddressController, clientRepository *repositories.ClientRepository) {
	r.DELETE("/address/:id", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
		}
		var addressId types.Address

		if err := c.ShouldBindUri(&addressId); err != nil {
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

		// find address to check auth
		address, errFind := controller.GetAddress(&types.Address{
			ID: addressId.ID,
		})

		if errFind != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "address not exists"})
			return
		}

		// Check auth client its the same address client
		if address.Client != client.ID {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		addressDeleted, errDeleteAddress := controller.DeleteAddress(&types.Address{
			ID: addressId.ID,
		})

		if errDeleteAddress != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errDeleteAddress.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, addressDeleted)
		}
	})
}
