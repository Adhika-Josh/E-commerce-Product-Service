package admin_controller

import (
	"e-commerce-product-service/app"
	"e-commerce-product-service/service/admin_service"
	"e-commerce-product-service/validator"
	"e-commerce-product-service/validator/admin_validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

var a admin_service.AdminServiceInterface = admin_service.AdminServiceImpl{
	DB: app.ConnectDB(),
}

func AddAdmins(c *gin.Context) {
	req, err := admin_validator.ValidateCreateAdmin(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := a.CreateAdmin(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}

func LoginAdmin(c *gin.Context) {
	req, err := admin_validator.ValidateLoginAdmin(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := a.LoginAdmin(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}
func DeleteAdmin(c *gin.Context) {
	req, err := admin_validator.ValidateDeleteAdmin(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := a.DeleteAdmin(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}
func UpdateAdmin(c *gin.Context) {
	req, err := admin_validator.ValidateUpdateAdmin(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := a.UpdateAdmin(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}
