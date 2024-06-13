package product_controller

import (
	"e-commerce-product-service/app"

	"e-commerce-product-service/service/product_service"
	"e-commerce-product-service/validator"
	"e-commerce-product-service/validator/product_validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

var p product_service.ProductServiceInterface = product_service.ProductServiceImpl{
	DB: app.ConnectDB(),
}

func AddProducts(c *gin.Context) {
	req, err := product_validator.ValidateAddProduct(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := p.AddProduct(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}

func GetAllProducts(c *gin.Context) {
	apiRes, err := p.GetAllProducts(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}

func UpdateProduct(c *gin.Context) {
	req, err := product_validator.ValidateUpdateProduct(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := p.UpdateProduct(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}

func GetProductsById(c *gin.Context) {
	id := c.Param("id")
	apiRes, err := p.GetProductByID(c, id)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)

}
