package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lavanderia.pro/internal/lavanderia/controllers"
)

func NewLaundryRouter(r *gin.Engine, c *controllers.LaundryController) {
	laundries, err := c.Laundries()
	r.GET("/laundries", func(c *gin.Context) {
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, laundries)
		}
	})
}
