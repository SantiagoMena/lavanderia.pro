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

		orderObject.Client.Auth = authId.(string)

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
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		// Handle Delete
		orderDeleted, errDelete := orderController.DeleteOrder(&types.Order{ID: order.ID})

		if errDelete != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "error on delete order"})
			return
		}

		c.IndentedJSON(http.StatusOK, orderDeleted)
	})
}

func NewPostAcceptOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business-order/accept/:id", func(c *gin.Context) {
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
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		// Handle Delete
		orderAccepted, errAccept := orderController.AcceptOrder(&types.Order{ID: order.ID})

		if errAccept != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "error on accept order"})
			return
		}

		c.IndentedJSON(http.StatusOK, orderAccepted)
	})
}

func NewRejectOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business-order/reject/:id", func(c *gin.Context) {
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
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		// Handle Delete
		orderRejected, errReject := orderController.RejectOrder(&types.Order{ID: order.ID})

		if errReject != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "error on reject order"})
			return
		}

		c.IndentedJSON(http.StatusOK, orderRejected)
	})
}

func NewAssignPickUpOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business-order/assign-pickup/:id", func(c *gin.Context) {
		authId := c.MustGet("auth")

		var orderId types.Order
		if err := c.ShouldBindUri(&orderId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		var delivery types.Delivery
		// Call BindJSON to bind the received JSON to
		// delivery.
		if errDeliveryJson := c.ShouldBindBodyWith(&delivery, binding.JSON); errDeliveryJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errDeliveryJson})
			return
		}

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		// Handle Controller
		order, err := orderController.GetOrder(&orderId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		order.Delivery = delivery

		// Handle Delete
		orderAssignedPickup, errReject := orderController.AssignPickUpOrder(&order)

		if errReject != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "error on assign pickup order"})
			return
		}

		c.IndentedJSON(http.StatusOK, orderAssignedPickup)
	})
}

func NewPickUpOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business-order/pickup-client/:id", func(c *gin.Context) {
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
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		orderPickupClient, errPickUp := orderController.PickUpClientOrder(&order)

		if errPickUp != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "error on pickup order"})
			return
		}

		c.IndentedJSON(http.StatusOK, orderPickupClient)
	})
}

func NewProcessOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business-order/process/:id", func(c *gin.Context) {
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
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		orderProcess, errProcess := orderController.ProcessOrder(&order)

		if errProcess != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "error on process order"})
			return
		}

		c.IndentedJSON(http.StatusOK, orderProcess)
	})
}

func NewFinishOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business-order/finish/:id", func(c *gin.Context) {
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
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		orderFinished, errFinish := orderController.FinishOrder(&order)

		if errFinish != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "error on finish order"})
			return
		}

		c.IndentedJSON(http.StatusOK, orderFinished)
	})
}

func NewAssignDeliveryOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business-order/assign-delivery/:id", func(c *gin.Context) {
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
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		orderFinished, errFinish := orderController.AssignDeliveryOrder(&order)

		if errFinish != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "error on assign delivery order"})
			return
		}

		c.IndentedJSON(http.StatusOK, orderFinished)
	})
}

func NewPickUpBusinessOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business-order/pickup-business/:id", func(c *gin.Context) {
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
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		orderPicketUpBusiness, errPickUpBusiness := orderController.PickUpBusinessOrder(&order)

		if errPickUpBusiness != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "error on pickup business order"})
			return
		}

		c.IndentedJSON(http.StatusOK, orderPicketUpBusiness)
	})
}

func NewDeliveryClientOrderRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business-order/delivery-client/:id", func(c *gin.Context) {
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
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if order.Client.Auth != authId.(string) {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		orderDeliveredClient, errDeliveryClient := orderController.DeliveryClientOrder(&order)

		if errDeliveryClient != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "error on delivery client order"})
			return
		}

		c.IndentedJSON(http.StatusOK, orderDeliveredClient)
	})
}
