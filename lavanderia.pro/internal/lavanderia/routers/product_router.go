package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
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

		var productObject types.Product
		// Call BindJSON to bind the received JSON to
		// newProduct.
		if err := c.BindJSON(&productObject); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
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

func NewDeleteProductRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	businessController *controllers.BusinessController,
) {
	r.DELETE("/product/:id", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
		}
		var productId types.Product

		if err := c.ShouldBindUri(&productId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		// Find Business and Check Auth
		productFound, errFind := productController.GetProduct(&productId)
		if errFind != nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": errFind.Error()})
			return
		}

		var unmarshalObjecth types.Product

		// convert m to s
		marshalObject, _ := bson.Marshal(productFound)
		bson.Unmarshal(marshalObject, &unmarshalObjecth)

		// Find Business and Check Auth
		businessFound, errFindBusiness := businessController.GetBusiness(&types.Business{
			ID: unmarshalObjecth.Business,
		})

		if errFindBusiness != nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": errFindBusiness.Error()})
			return
		}

		if string(businessFound.Auth) != authId {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
			return
		}

		products, err := productController.DeleteProduct(&productId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, products)
		}

	})
}

func NewGetProductRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
) {
	r.GET("/product/:id", func(c *gin.Context) {
		var productId types.Product

		if err := c.ShouldBindUri(&productId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		// Find Business and Check Auth
		productFound, errFind := productController.GetProduct(&productId)
		if errFind != nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": errFind.Error()})
			return
		}

		if errFind != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": errFind.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, productFound)
		}

	})
}

func NewUpdateProductRouter(
	r *gin.Engine,
	productController *controllers.ProductController,
	businessController *controllers.BusinessController,
) {
	r.PUT("/product/:id", func(c *gin.Context) {
		authId := c.MustGet("auth")

		if authId == nil {
			c.JSON(http.StatusForbidden, gin.H{"msg": "permissions denied"})
		}

		var productId types.Product

		if err := c.ShouldBindUri(&productId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		// Find Product and Check Auth
		productFound, errFindProduct := productController.GetProduct(&productId)
		if errFindProduct != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": errFindProduct.Error()})
			return
		}

		businessFound, errFindBusiness := businessController.GetBusiness(&types.Business{
			ID: productFound.Business,
		})

		if errFindBusiness != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": errFindBusiness.Error()})
			return
		}

		if string(businessFound.Auth) != authId {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "permissions denied"})
			return
		}

		var productObject types.Product
		// Call BindJSON to bind the received JSON to
		// newProduct.
		if err := c.BindJSON(&productObject); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		if string(businessFound.Auth) != authId {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "permissions denied"})
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
