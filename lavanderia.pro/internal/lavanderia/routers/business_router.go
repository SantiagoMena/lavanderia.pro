package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/controllers"
)

func NewGetAllBusinessRouter(r *gin.Engine, controller *controllers.BusinessController) {
	r.GET("/business", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		business, err := controller.GetAllBusinessByAuth(authId.(string))

		if business == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		} else {
			c.JSON(http.StatusOK, business)
		}
	})
}

func NewPostBusinessRouter(r *gin.Engine, controller *controllers.BusinessController) {
	r.POST("/business", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		var newBusiness types.Business

		// Call BindJSON to bind the received JSON to
		// newBusiness.
		if err := c.BindJSON(&newBusiness); err != nil {
			return
		}

		newBusiness.Auth = authId.(string)

		// Handle Controller
		business, err := controller.PostBusiness(&newBusiness)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		} else {
			c.IndentedJSON(http.StatusCreated, business)
		}

	})
}

// type BusinessId struct {
// 	ID string `uri:"id" binding:"required,uuid"`
// }

func NewDeleteBusinessRouter(r *gin.Engine, controller *controllers.BusinessController) {
	r.DELETE("/business/:id", func(c *gin.Context) {
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
		businessFound, errFind := controller.GetBusiness(&businessId)
		if errFind != nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": errFind})
			return
		}

		if string(businessFound.Auth) != authId {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		// Handle Controller
		deletedBusiness, err := controller.DeleteBusiness(&businessFound)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		} else {
			c.IndentedJSON(http.StatusCreated, deletedBusiness)
		}

	})
}

func NewUpdateBusinessRouter(r *gin.Engine, controller *controllers.BusinessController) {
	r.PUT("/business/:id", func(c *gin.Context) {
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
		businessFound, errFind := controller.GetBusiness(&businessId)
		if errFind != nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": errFind})
			return
		}

		if string(businessFound.Auth) != authId {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		var business types.Business
		if errJson := c.ShouldBindBodyWith(&business, binding.JSON); errJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errJson})
			return
		}

		// Handle Controller
		updatedBusiness, err := controller.UpdateBusiness(&types.Business{
			ID:       businessFound.ID,
			Name:     business.Name,
			Position: business.Position,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.IndentedJSON(http.StatusCreated, updatedBusiness)
		}

	})
}

func NewGetBusinessRouter(r *gin.Engine, controller *controllers.BusinessController) {
	r.GET("/business/:id", func(c *gin.Context) {
		var business types.Business

		if err := c.ShouldBindUri(&business); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		// Handle Controller
		businessDb, err := controller.GetBusiness(&business)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, businessDb)
		}

	})
}

func NewPostRegisterBusinessRouter(r *gin.Engine, controller *controllers.AuthController) {
	r.POST("/business/register", func(c *gin.Context) {

		var newBusiness types.Business
		var newAuth types.Auth

		// Call BindJSON to bind the received JSON to
		// newBusiness.
		if errBusinessJson := c.ShouldBindBodyWith(&newBusiness, binding.JSON); errBusinessJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errBusinessJson})
			return
		}

		// Call BindJSON to bind the received JSON to
		// newAuth.
		if errAuthJson := c.ShouldBindBodyWith(&newAuth, binding.JSON); errAuthJson != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": errAuthJson})
			return
		}

		// Handle Controller
		business, errRegister := controller.RegisterBusiness(&newAuth, &newBusiness)

		if errRegister != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": errRegister.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, business)
		}

	})
}
func NewPostRegisterBusinessDeliveryRouter(r *gin.Engine, controller *controllers.BusinessController) {
	r.POST("/business/:id/delivery", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		var newDelivery types.Delivery
		var businessId types.Business
		var newAuth types.Auth

		// Call BindJSON to bind the received JSON to
		// businessId.
		if err := c.ShouldBindUri(&businessId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

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
		delivery, errRegister := controller.RegisterBusinessDelivery(&newAuth, &businessId, &newDelivery)

		fmt.Println(businessId)
		if errRegister != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": errRegister.Error()})
		} else {
			c.IndentedJSON(http.StatusCreated, delivery)
		}

	})
}
