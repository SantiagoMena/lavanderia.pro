package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getIndex)

	router.Run("0.0.0.0:1002")
}

func getIndex(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}
