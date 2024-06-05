package product_service

import (
	"e-commerce-product-service/model"
	"e-commerce-product-service/model/entity"
	"e-commerce-product-service/utils"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	DB *gorm.DB
}

func (p ProductServiceImpl) AddProduct(c *gin.Context, req model.AddProductsRequest) (model.AddProductResponse, model.Errors) {
	var res model.AddProductResponse
	var custErr model.Errors
	var existingProduct entity.ProductDetails
	if req.ItemPID != "" {
		if err := p.DB.Where("item_pid = ?", req.ItemPID).First(&existingProduct).Error; err == nil {
			return res, model.Errors{
				Error: "Item already exists with the same Item_PID",
				Type:  "duplicate_entry_error",
			}
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{
				Error: err.Error(),
				Type:  "internal_server_error",
			}
		}
	} else {
		if err := p.DB.Where("item_name = ? AND item_category = ? AND item_price = ?", req.ItemName, req.ItemCategory, req.ItemPrice).First(&existingProduct).Error; err == nil {
			return res, model.Errors{
				Error: "Item already exists with the same name, category, and price",
				Type:  "duplicate_entry_error",
			}
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{
				Error: err.Error(),
				Type:  "internal_server_error",
			}
		}
		newProduct := entity.ProductDetails{
			ItemPID:      "PD_" + utils.GenerateRandString(6),
			ItemName:     req.ItemName,
			ItemCategory: req.ItemCategory,
			ItemQuantity: req.ItemQuantity,
			ItemPrice:    req.ItemPrice,
			CreatedAt:    time.Now(),
		}
		if err := p.DB.Create(&newProduct).Error; err != nil {
			return res, model.Errors{
				Error: "Failed to add product",
				Type:  "internal_server_error",
			}
		}
		return model.AddProductResponse{
			Status:  http.StatusOK,
			Message: "Added Products Successfully",
			ItemPID: newProduct.ItemPID,
		}, model.Errors{}
	}
	return res, custErr
}

func (p ProductServiceImpl) GetAllProducts(c *gin.Context) ([]model.GetAllProductsResponse, model.Errors) {
	var products []entity.ProductDetails
	if err := p.DB.Find(&products).Error; err != nil {
		return nil, model.Errors{
			Error: err.Error(),
			Type:  "internal_server_error",
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.Errors{
			Error: "no record found in the table",
			Type:  "no_record_found",
		}
	}

	var response []model.GetAllProductsResponse
	for _, product := range products {
		response = append(response, model.GetAllProductsResponse{
			ItemPID:      product.ItemPID,
			ItemName:     product.ItemName,
			ItemCategory: product.ItemCategory,
			ItemQuantity: product.ItemQuantity,
			ItemPrice:    product.ItemPrice,
		})
	}
	return response, model.Errors{}
}

func (p ProductServiceImpl) UpdateProduct(c *gin.Context, req model.UpdateProductRequest) (model.UpdateProductResponse, model.Errors) {
	var res model.UpdateProductResponse
	var custErr model.Errors
	var existingProduct entity.ProductDetails

	if err := p.DB.Where("item_pid = ?", req.ItemPID).First(&existingProduct).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{
				Error: "No record found for the given Item_PID",
				Type:  "record_not_found",
			}
		}
	}
	existingProduct.ItemQuantity = req.ItemQuantity
	existingProduct.ItemPrice = req.ItemPrice
	existingProduct.UpdatedAt = time.Now()
	if err := p.DB.Save(&existingProduct).Error; err != nil {
		return res, model.Errors{
			Error: "Failed to update product",
			Type:  "internal_server_error",
		}
	}
	res.Status = http.StatusOK
	res.Message = "Product updated successfully"
	res.ItemPID = existingProduct.ItemPID
	return res, custErr
}
