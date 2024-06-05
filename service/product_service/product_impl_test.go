package product_service

import (
	"e-commerce-product-service/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var c *gin.Context

func TestAddProductsWithNilItemPID(t *testing.T) {
	req := model.AddProductsRequest{
		ItemName:     "Charger",
		ItemCategory: "Budget",
		ItemQuantity: "10",
		ItemPrice:    "5",
		ItemPID:      "",
	}

	res, err := ProductServiceImpl{}.AddProduct(c, req)
	expectedRes := model.AddProductResponse{
		Status:  200,
		Message: "Added Products Successfully",
		ItemPID: res.ItemPID,
	}
	assert.NoError(t, err.Message)
	assert.NotEmpty(t, res)
	assert.Equal(t, expectedRes, res)
}

func TestAddProductsWithItemPID(t *testing.T) {
	req := model.AddProductsRequest{
		ItemName:     "Charger",
		ItemCategory: "Budget",
		ItemQuantity: "10",
		ItemPrice:    "5",
		ItemPID:      "PD_123ABH",
	}

	_, err := ProductServiceImpl{}.AddProduct(c, req)
	expectedRes := model.Errors{
		Error: "Item already exists with the same Item_PID",
		Type:  "duplicate_entry_error",
	}
	assert.NoError(t, err.Message)
	assert.Equal(t, expectedRes, err)
}

func TestAddProductsWithsameItem(t *testing.T) {
	req := model.AddProductsRequest{
		ItemName:     "Charger",
		ItemCategory: "Budget",
		ItemQuantity: "10",
		ItemPrice:    "5",
		ItemPID:      "",
	}

	_, err := ProductServiceImpl{}.AddProduct(c, req)
	expectedRes := model.Errors{
		Error: "Item already exists with the same name, category, and price",
		Type:  "duplicate_entry_error",
	}
	assert.NoError(t, err.Message)
	assert.Equal(t, expectedRes, err)
}

func TestGetAllProducts(t *testing.T) {
	expectedRes := []model.GetAllProductsResponse{
		{
			ItemPID:      "PD_123ABH",
			ItemName:     "Pen",
			ItemCategory: "Budget",
			ItemQuantity: "10",
			ItemPrice:    "10.00",
		},
		{
			ItemPID:      "PD_56ADR",
			ItemName:     "Book",
			ItemCategory: "Regular",
			ItemQuantity: "20",
			ItemPrice:    "30.00",
		},
		{
			ItemPID:      "PD_RT12QW",
			ItemName:     "Bottle",
			ItemCategory: "Regular",
			ItemQuantity: "30",
			ItemPrice:    "100.00",
		},
	}

	res, err := ProductServiceImpl{}.GetAllProducts(c)
	assert.NoError(t, err.Message)
	assert.Equal(t, expectedRes, res)
}

func TestUpdateProductWithValidItemPID(t *testing.T) {
	expectedRes := model.UpdateProductResponse{
		Status:  200,
		Message: "Product updated successfully",
		ItemPID: "PD_123ABH",
	}
	req := model.UpdateProductRequest{
		ItemPID:      "PD_123ABH",
		ItemQuantity: "200",
		ItemPrice:    "1.00",
	}
	res, err := ProductServiceImpl{}.UpdateProduct(c, req)
	assert.NoError(t, err.Message)
	assert.Equal(t, expectedRes, res)
}

func TestUpdateProductWithInValidItemPID(t *testing.T) {
	expectedRes := model.Errors{
		Error: "No record found for the given Item_PID",
		Type:  "record_not_found",
	}
	req := model.UpdateProductRequest{
		ItemPID:      "PD_120ABH",
		ItemQuantity: "200",
		ItemPrice:    "1.00",
	}
	_, err := ProductServiceImpl{}.UpdateProduct(c, req)
	assert.Error(t, err.Message)
	assert.Equal(t, expectedRes, err)
}
