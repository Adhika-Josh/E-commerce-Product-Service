package admin_validator

import (
	"e-commerce-product-service/model"
	"e-commerce-product-service/validator"
	"regexp"

	"github.com/gin-gonic/gin"
)

func ValidateCreateAdmin(c *gin.Context) (req model.CreateAdminRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	if req.Name == "" || req.PhoneNo == "" || req.Dob == "" || req.Email == "" || req.UserName == "" || req.Password == "" {
		return req, model.Errors{
			Error: "Missing mandatory fields",
			Type:  "validation_error",
		}
	}
	// Regular expressions
	phoneRegex := `^\d{10}$`
	emailRegex := `^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	dobRegex := `\d{4}-\d{2}-\d{2}`
	if matched, _ := regexp.MatchString(phoneRegex, req.PhoneNo); !matched {
		return req, model.Errors{
			Error: "Invalid phone number format",
			Type:  "validation_error",
		}
	}
	if matched, _ := regexp.MatchString(emailRegex, req.Email); !matched {
		return req, model.Errors{
			Error: "Invalid email format",
			Type:  "validation_error",
		}
	}
	if matched, _ := regexp.MatchString(dobRegex, req.Dob); !matched {
		return req, model.Errors{
			Error: "Invalid date of birth format",
			Type:  "validation_error",
		}
	}
	return req, custErr
}
func ValidateLoginAdmin(c *gin.Context) (req model.LoginAdminRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	return req, custErr
}
func ValidateDeleteAdmin(c *gin.Context) (req model.DeleteAdminRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	return req, custErr
}
func ValidateUpdateAdmin(c *gin.Context) (req model.UpdateAdminRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	return req, custErr
}
