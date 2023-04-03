package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/controllers"
)

func NewGetLaundriesRouter(r *gin.Engine, controller *controllers.LaundryController) {
	r.GET("/laundries", func(c *gin.Context) {
		laundries, err := controller.GetLaundries()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, laundries)
		}
	})
}

func NewPostLaundrysRouter(r *gin.Engine, controller *controllers.LaundryController) {
	r.POST("/laundry", func(c *gin.Context) {

		var newLaundry types.Laundry

		// Call BindJSON to bind the received JSON to
		// newLaundry.
		if err := c.BindJSON(&newLaundry); err != nil {
			return
		}

		// Handle Controller
		laundry, err := controller.PostLaundry(&newLaundry)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.IndentedJSON(http.StatusCreated, laundry)
		}

	})
}

// type LaundryId struct {
// 	ID string `uri:"id" binding:"required,uuid"`
// }

func NewDeleteLaundrysRouter(r *gin.Engine, controller *controllers.LaundryController) {
	r.DELETE("/laundry/:id", func(c *gin.Context) {
		var laundry types.Laundry

		if err := c.ShouldBindUri(&laundry); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		// Handle Controller
		deletedLaundry, err := controller.DeleteLaundry(&laundry)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.IndentedJSON(http.StatusCreated, deletedLaundry)
		}

	})
}

func NewUpdateLaundrysRouter(r *gin.Engine, controller *controllers.LaundryController) {
	r.PUT("/laundry/:id", func(c *gin.Context) {
		var laundry types.Laundry

		if err := c.ShouldBindUri(&laundry); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		if errJson := c.ShouldBindBodyWith(&laundry, binding.JSON); errJson != nil {
			c.JSON(400, gin.H{"msg": errJson})
			return
		}

		// Handle Controller
		updatedLaundry, err := controller.UpdateLaundry(&laundry)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.IndentedJSON(http.StatusCreated, updatedLaundry)
		}

	})
}
