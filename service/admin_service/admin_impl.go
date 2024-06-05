package admin_service

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

type AdminServiceImpl struct {
	DB *gorm.DB
}

func (a AdminServiceImpl) CreateAdmin(c *gin.Context, req model.CreateAdminRequest) (model.CreateAdminResponse, model.Errors) {
	var res model.CreateAdminResponse
	var custErr model.Errors

	var existingAdmin entity.AdminDetails
	if err := a.DB.Where("email = ?", req.Email).First(&existingAdmin).Error; err == nil {
		return res, model.Errors{
			Error: "Admin with same email already exists",
			Type:  "duplicate_entry_error",
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, model.Errors{
			Error: err.Error(),
			Type:  "internal_server_error",
		}
	}
	newAdmin := entity.AdminDetails{
		AdminPID:  "AD_" + utils.GenerateRandString(8),
		Name:      req.Name,
		PhoneNo:   req.PhoneNo,
		Dob:       req.Dob,
		Email:     req.Email,
		Password:  req.Password,
		UpdatedAt: time.Now(),
	}
	if err := a.DB.Create(&newAdmin).Error; err != nil {
		return res, model.Errors{
			Error: "Failed to create admin",
			Type:  "internal_server_error",
		}
	}
	res.Status = http.StatusOK
	res.Message = "Admin created successfully"
	res.AdminPID = newAdmin.AdminPID

	return res, custErr
}

func (a AdminServiceImpl) LoginAdmin(c *gin.Context, req model.LoginAdminRequest) (model.CreateAdminResponse, model.Errors) {
	var res model.CreateAdminResponse
	var custErr model.Errors

	var admin entity.AdminDetails
	if err := a.DB.Where("user_name = ?", req.UserName).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{
				Error: "Invalid username or password",
				Type:  "invalid_credentials_error",
			}
		}
		return res, model.Errors{
			Error: err.Error(),
			Type:  "internal_server_error",
		}
	}
	if admin.UserName != req.UserName || admin.Password != req.Password {
		return res, model.Errors{
			Error: "Invalid username or password",
			Type:  "invalid_credentials_error",
		}
	}
	res.Status = http.StatusOK
	res.Message = "Login successful"
	res.AdminPID = admin.AdminPID

	return res, custErr
}
