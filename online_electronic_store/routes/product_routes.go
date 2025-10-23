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
	r.GET("/products", productTable.GetProducts)                    
	r.GET("/product/:id", productTable.GetProductById)              
	r.GET("/products/:category", productTable.GetProductByCategory) 
	r.POST("/product/create", middleware.AuthMiddleware(), productTable.CreateProduct)
	r.PUT("/product/:product_id/update", middleware.AuthMiddleware(), productTable.UpdateProductOnId)
}
