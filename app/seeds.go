package app

import (
	"e-commerce-product-service/model/entity"
	"time"

	"gorm.io/gorm"
)

func seedProductDetails(db *gorm.DB) error {
	products := []entity.ProductDetails{
		{ItemPID: "PD_123ABH", ItemName: "Pen", ItemCategory: "Budget", ItemQuantity: "10", ItemPrice: " 10.00", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ItemPID: "PD_56ADRT", ItemName: "Book", ItemCategory: "Regular", ItemQuantity: "20", ItemPrice: " 30.00", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ItemPID: "PD_RT12QW", ItemName: "Bottle", ItemCategory: "Regular", ItemQuantity: "30", ItemPrice: " 100.00", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, p := range products {
		err := db.Create(&p).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func seedAdminDetails(db *gorm.DB) error {
	admins := []entity.AdminDetails{
		{AdminPID: "AD_123456", Name: "John Doe", PhoneNo: "1234567890", Dob: "1990-01-01", Email: "john@example.com", UserName: "John123", Password: "John@123", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{AdminPID: "AD_789012", Name: "Jane Smith", PhoneNo: "9876543210", Dob: "1995-05-15", Email: "jane@example.com", UserName: "Jane123", Password: "Jane@123", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, admin := range admins {
		err := db.Create(&admin).Error
		if err != nil {
			return err
		}
	}
	return nil
}
