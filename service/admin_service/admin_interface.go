package admin_service

import (
	"e-commerce-product-service/model"

	"github.com/gin-gonic/gin"
)

type AdminServiceInterface interface {
	CreateAdmin(c *gin.Context, req model.CreateAdminRequest) (model.CreateAdminResponse, model.Errors)
	LoginAdmin(c *gin.Context, req model.LoginAdminRequest) (model.CreateAdminResponse, model.Errors)
}
