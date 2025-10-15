package routes

import (
	"online_electronic_store/handlers"
	"online_electronic_store/interfaces"
	"online_electronic_store/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeliveryRoutes(r *gin.Engine, db *gorm.DB) {
	var deliveryTable interfaces.Deliveries = handlers.NewDeliveryTable(db)
	r.GET("/delivery/:delivery_id", middleware.AuthMiddleware(), deliveryTable.GetDeliveryDetailsByOrderID)
	r.GET("/delivery", middleware.AuthMiddleware(), deliveryTable.GetDeliveryDetailsList)        
	r.POST("/delivery/create/:order_ID", middleware.AuthMiddleware(), deliveryTable.AddDelivery) 
}
