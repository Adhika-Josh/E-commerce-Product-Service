package product_service

import (
	"e-commerce-product-service/model"

	"github.com/gin-gonic/gin"
)

type ProductServiceInterface interface {
	AddProduct(c *gin.Context, req model.AddProductsRequest) (model.AddProductResponse, model.Errors)
	GetAllProducts(c *gin.Context) ([]model.GetAllProductsResponse, model.Errors)
	UpdateProduct(c *gin.Context, req model.UpdateProductRequest) (model.UpdateProductResponse, model.Errors)
	GetProductByID(c *gin.Context, id string) (model.GetProductsByIDResponse, model.Errors)
}
