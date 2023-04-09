package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/controllers"
)

func NewPostOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business/:id/order", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		var businessId types.Business

		if err := c.ShouldBindUri(&businessId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		// Find Business and Check Auth
		businessFound, errFind := businessController.GetBusiness(&businessId)
		if errFind != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": errFind.Error()})
			return
		}

		if string(businessFound.Auth) != authId {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "permissions denied"})
			return
		}

		// Call BindJSON to bind the received JSON to
		// client.
		var client types.Client
		if errClientJson := c.ShouldBindBodyWith(&client, binding.JSON); errClientJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errClientJson})
			return
		}

		// Call BindJSON to bind the received JSON to
		// address.
		var address types.Address
		if errAddressJson := c.ShouldBindBodyWith(&address, binding.JSON); errAddressJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errAddressJson})
			return
		}

		// Call BindJSON to bind the received JSON to
		// productList.
		var productList []types.OrderProduct
		if errProductJson := c.ShouldBindBodyWith(&productList, binding.JSON); errProductJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errProductJson})
			return
		}

		// Handle Controller
		order, err := orderController.PostOrder(&types.Order{
			Client:   client,
			Business: businessFound,
			Address:  address,
			Products: productList,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		} else {
			c.IndentedJSON(http.StatusCreated, order)
		}

	})
}
