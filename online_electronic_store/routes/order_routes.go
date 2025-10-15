package routes

import (
	"online_electronic_store/handlers"
	"online_electronic_store/interfaces"
	"online_electronic_store/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRoutes(r *gin.Engine, db *gorm.DB) {
	var orderTable interfaces.Orders = handlers.NewOrderTable(db)
	r.GET("/orders", middleware.AuthMiddleware(), orderTable.GetOrders)             //all the orders of the user
	r.POST("/orders/create", middleware.AuthMiddleware(), orderTable.MakeOrder)     //create the new order for the user
	r.PUT("/orders/:order_id", middleware.AuthMiddleware(), orderTable.UpdateOrder) // updating the order
}

//
