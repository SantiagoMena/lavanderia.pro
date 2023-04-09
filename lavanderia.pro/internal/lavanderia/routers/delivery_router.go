package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/controllers"
)

func NewPostRegisterDeliveryRouter(r *gin.Engine, controller *controllers.AuthController) {
	r.POST("/delivery/register", func(c *gin.Context) {

		var newDelivery types.Delivery
		var newAuth types.Auth

		// Call BindJSON to bind the received JSON to
		// newDelivery.
		if errDeliveryJson := c.ShouldBindBodyWith(&newDelivery, binding.JSON); errDeliveryJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errDeliveryJson})
			return
		}

		// Call BindJSON to bind the received JSON to
		// newAuth.
		if errAuthJson := c.ShouldBindBodyWith(&newAuth, binding.JSON); errAuthJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errAuthJson})
			return
		}

		// Handle Controller
		delivery, errRegister := controller.RegisterDelivery(&newAuth, &newDelivery)

		if errRegister != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errRegister.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, delivery)
		}

	})
}

func NewPostDeliveryRouter(r *gin.Engine, controller *controllers.DeliveryController) {
	r.POST("/delivery/profile", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
		}

		var delivery types.Delivery

		if err := c.BindJSON(&delivery); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		delivery.Auth = authId.(string)

		// Handle Controller
		deliveryPosted, errPost := controller.PostDelivery(&delivery)

		if errPost != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errPost.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, deliveryPosted)
		}

	})
}
