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
	r.POST("/business-order/:id", func(c *gin.Context) {
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
		var orderObject types.Order
		if errOrderJson := c.ShouldBindBodyWith(&orderObject, binding.JSON); errOrderJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errOrderJson.Error() + "1"})
			return
		}

		// Handle Controller
		order, err := orderController.PostOrder(&types.Order{
			Client:   orderObject.Client,
			Business: businessFound,
			Address:  orderObject.Address,
			Products: orderObject.Products,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		} else {
			c.IndentedJSON(http.StatusCreated, order)
		}

	})
}

func NewGetOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.GET("/business-order/:id", func(c *gin.Context) {
		authId := c.MustGet("auth")

		var orderId types.Order
		if err := c.ShouldBindUri(&orderId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		// Handle Controller
		order, err := orderController.GetOrder(&orderId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		} else {
			c.IndentedJSON(http.StatusOK, order)
		}

	})
}

func NewDeleteOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.DELETE("/business-order/:id", func(c *gin.Context) {
		authId := c.MustGet("auth")

		var orderId types.Order
		if err := c.ShouldBindUri(&orderId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		// Handle Controller
		order, err := orderController.GetOrder(&orderId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
		}

		// Handle Delete
		orderDeleted, errDelete := orderController.DeleteOrder(&types.Order{ID: order.ID})

		if errDelete != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		}

		c.IndentedJSON(http.StatusOK, orderDeleted)
	})
}
