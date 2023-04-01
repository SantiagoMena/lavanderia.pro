package routers

import (
	// "log"

	// "github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"lavanderia.pro/internal/lavanderia/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.GET("/laundries", controllers.Laundries)

	return r
}
