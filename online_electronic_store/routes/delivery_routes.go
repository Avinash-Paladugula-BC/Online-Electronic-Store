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
	//done
	r.GET("/delivery/:delivery_id", middleware.AuthMiddleware(), deliveryTable.GetDeliveryDetailsByOrderID) // the specific order of the customer delivery details needs to be returned
	r.GET("/delivery", middleware.AuthMiddleware(), deliveryTable.GetDeliveryDetailsList)                            // get the list of all delivery updates of a user
	r.POST("/delivery/create/:order_ID", middleware.AuthMiddleware(), deliveryTable.AddDelivery)                     // when a person makes a new order the delivery is created
	// r.PUT("/delivery/:delivery_id/update", controller.DeliveryUpdate) // not required
}
