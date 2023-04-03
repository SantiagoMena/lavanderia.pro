package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
