package app

import (
	"e-commerce-product-service/model/entity"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("./product_service.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.ProductDetails{}, &entity.AdminDetails{})
	if err != nil {
		panic("Database migration failed")
	}
	err = seedDatabase(db)
	if err != nil {
		log.Printf("Error occurred while seeding products database: %v", err.Error())
		return nil
	}
	return db
}

func seedDatabase(db *gorm.DB) error {
	if err := seedProductDetails(db); err != nil {
		return err
	}
	if err := seedAdminDetails(db); err != nil {
		return err
	}
	return nil
}
