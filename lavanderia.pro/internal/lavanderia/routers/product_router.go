package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/controllers"
)

func NewPostProductRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	businessController *controllers.BusinessController,
) {
	r.POST("/business/:id/product", func(c *gin.Context) {
		authId := c.MustGet("auth")

		var businessId types.Business

		if err := c.ShouldBindUri(&businessId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		// Find Business and Check Auth
		businessFound, errFind := businessController.GetBusiness(&businessId)
		if errFind != nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": errFind.Error()})
			return
		}

		if string(businessFound.Auth) != authId {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		var productObject types.Product
		// Call BindJSON to bind the received JSON to
		// newProduct.
		if err := c.BindJSON(&productObject); err != nil {
			return
		}

		productObject.Business = businessFound.ID

		// Handle Controller
		product, err := productController.PostProduct(&productObject)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		} else {
			c.IndentedJSON(http.StatusCreated, product)
		}

	})
}

func NewGetProductsByBusinessRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	businessController *controllers.BusinessController,
) {
	r.GET("/business/:id/products", func(c *gin.Context) {
		var businessId types.Business

		if err := c.ShouldBindUri(&businessId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		products, err := productController.GetAllProductsByBusiness(string(businessId.ID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		} else {
			c.IndentedJSON(http.StatusOK, products)
		}

	})
}
