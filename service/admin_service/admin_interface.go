package admin_service

import (
	"e-commerce-product-service/model"

	"github.com/gin-gonic/gin"
)

type AdminServiceInterface interface {
	CreateAdmin(c *gin.Context, req model.CreateAdminRequest) (model.CreateAdminResponse, model.Errors)
	LoginAdmin(c *gin.Context, req model.LoginAdminRequest) (model.CreateAdminResponse, model.Errors)
	DeleteAdmin(c *gin.Context, req model.DeleteAdminRequest) (model.DeleteAdminResponse, model.Errors)
	UpdateAdmin(c *gin.Context, req model.UpdateAdminRequest) (model.UpdateAdminResponse, model.Errors)
}
