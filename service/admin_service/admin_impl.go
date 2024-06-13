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
		AdminPID:  "AD_" + utils.GenerateRandString(6),
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
func (a AdminServiceImpl) DeleteAdmin(c *gin.Context, req model.DeleteAdminRequest) (model.DeleteAdminResponse, model.Errors) {
	var user entity.AdminDetails
	var res model.DeleteAdminResponse
	var custErr model.Errors
	if err := a.DB.Where("admin_pid= ?", req.AdminPID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{
				Error: "no_user_found",
				Type:  "invalid_credentials_error",
			}
		}
		return res, model.Errors{
			Error: err.Error(),
			Type:  "internal_server_error",
		}
	}
	if err := a.DB.Delete(&user).Error; err != nil {
		return res, model.Errors{
			Error: "Failed to delete user",
			Type:  "internal_server_error",
		}
	}
	res.Status = http.StatusOK
	res.Message = "User deleted successfully"
	res.AdminPID = user.AdminPID
	return res, custErr
}
func (a AdminServiceImpl) UpdateAdmin(c *gin.Context, req model.UpdateAdminRequest) (model.UpdateAdminResponse, model.Errors) {
	var admin entity.AdminDetails
	var res model.UpdateAdminResponse
	var custErr model.Errors

	if err := a.DB.Where("admin_pid = ?", req.AdminPID).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{
				Error: "No record found for the given Item_PID",
				Type:  "record_not_found",
			}
		}
	}

	admin.AdminPID = req.AdminPID
	admin.Name = req.Name
	admin.Dob = req.Dob
	admin.PhoneNo = req.PhoneNo
	admin.Email = req.Email
	admin.UserName = req.UserName
	admin.Password = req.Password
	admin.UpdatedAt = time.Now()
	if err := a.DB.Save(&admin).Error; err != nil {
		return res, model.Errors{
			Error: "Failed to update product",
			Type:  "internal_server_error",
		}
	}
	res.Status = http.StatusOK
	res.Message = "User Details updated successfully"
	res.AdminPID = admin.AdminPID
	return res, custErr
}
