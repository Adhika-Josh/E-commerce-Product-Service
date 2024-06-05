package product_validator

import (
	"e-commerce-product-service/model"
	"e-commerce-product-service/validator"

	"github.com/gin-gonic/gin"
)

func ValidateAddProduct(c *gin.Context) (req model.AddProductsRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	if req.ItemName == "" || req.ItemCategory == "" || req.ItemQuantity == "" || req.ItemPrice == "" {
		return req, model.Errors{
			Error: "Missing mandatory fields",
			Type:  "validation_error",
		}
	}
	return req, model.Errors{}
}
func ValidateUpdateProduct(c *gin.Context) (req model.UpdateProductRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	return req, model.Errors{}
}
