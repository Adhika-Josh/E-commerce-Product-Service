package product_validator

import (
	"bytes"
	"e-commerce-product-service/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var expectedRequest = model.AddProductsRequest{
	ItemName:     "Test Product",
	ItemCategory: "Test Category",
	ItemQuantity: "10",
	ItemPrice:    "100",
	ItemPID:      "PD_4545GH",
}

func TestValidateAddProduct(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	requestBody := []byte(`{
        "item_name": "Test Product",
        "item_category": "Test Category",
        "item_quantity": "10",
        "item_price": "100",
		"item_pid":"PD_4545GH"
    }`)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")
	req, err := ValidateAddProduct(c)
	assert.NoError(t, err.Message)
	assert.Equal(t, expectedRequest, req)
}
