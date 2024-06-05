package main

import (
	"e-commerce-product-service/app"
	admin_controller "e-commerce-product-service/controller/admin_controller"
	product_controller "e-commerce-product-service/controller/product_controller"
	"e-commerce-product-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	app.ConnectDB()
	r := gin.Default()
	r.Use(middleware.PanicHandling())
	r.Use(middleware.ValidateJSONMiddleware())
	admin := r.Group("/product-service/v1/admin")
	{
		admin.POST("/add", admin_controller.AddAdmins)
		admin.POST("/login", admin_controller.LoginAdmin)
	}
	v1 := r.Group("/product-service/v1")
	{
		v1.POST("/add-product", product_controller.AddProducts)
		v1.GET("/get-all-products", product_controller.GetAllProducts)
		v1.PATCH("/update-product", product_controller.UpdateProduct)
	}
	r.Run("localhost:8080")

}
