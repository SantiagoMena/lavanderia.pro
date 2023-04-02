package routers

import (
	// "log"

	// "github.com/joho/godotenv"

	"net/http"

	"github.com/gin-gonic/gin"
	"lavanderia.pro/internal/lavanderia/controllers"
)

func NewPingRouter(r *gin.Engine, c *controllers.PingController) {
	ping, err := c.Ping()
	r.GET("/ping", func(c *gin.Context) {
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, ping)
		}
	})
}
