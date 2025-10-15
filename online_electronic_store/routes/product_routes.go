package routes

import (
	"online_electronic_store/handlers"
	"online_electronic_store/interfaces"
	"online_electronic_store/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoutes(r *gin.Engine, db *gorm.DB) {
	var productTable interfaces.Products = handlers.NewProductTable(db)
	r.GET("/products", productTable.GetProducts)                    // get all the products
	r.GET("/product/:id", productTable.GetProductById)              // get details of the specific product
	r.GET("/products/:category", productTable.GetProductByCategory) // get all the products details of particular category
	//also need to add another middleware, to check if the user is admin
	r.POST("/product/create", middleware.AuthMiddleware(), productTable.CreateProduct)                // also need to include middleware inorder to check if the user is admin
	r.PUT("/product/:product_id/update", middleware.AuthMiddleware(), productTable.UpdateProductOnId) // change the details of the product && need to be an admin
}
