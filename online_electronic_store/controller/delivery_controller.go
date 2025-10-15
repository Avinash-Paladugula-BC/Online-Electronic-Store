package controller

// import (
// 	"fmt"
// 	"net/http"
// 	// "online_electronic_store/config"
// 	"online_electronic_store/interfaces"
// 	"online_electronic_store/models"
// 	"online_electronic_store/tables"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// // var Delivery_table tables.DeliveryTable
// var DeliveryTable interfaces.Online_Electronic_Store_DB[models.Delivery]         ///----------------------



// func InstantiateDeliveryTable(dbb *gorm.DB) {
// 	DeliveryTable = tables.NewDeliveryTable(dbb)
// 	fmt.Println("Instantiated delivery table", DeliveryTable)
// }
// ////////////////// done
// func ()GetDeliveryDetailsByProductID(c *gin.Context) {
// 	_, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
// 		return
// 	}
// 	// userID := userIDInterface.(uint)
// 	deliveryIDInterface := c.Param("product_id")
// 	// if !ok {
// 	// 	c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
// 	// }
// 	delivery_id, _ := strconv.Atoi(deliveryIDInterface)
// 	delivery, err := DeliveryTable.GetByID(c, uint(delivery_id))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch delivery details"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"Delivery ID": delivery.ID,
// 		"Delivery status": delivery.DeliveryStatus,
// 	})
// }

// func GetDeliveryDetailsList(c *gin.Context) {
// 	userIDInterface, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
// 		return
// 	}
// 	userID := userIDInterface.(uint)
// 	var deliveries []models.Delivery
// 	deliveries, err := DeliveryTable.GetAll(c)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch delivery details"})
// 		return
// 	}
// 	response := []gin.H{}
// 	for _, r := range deliveries{
// 		response=append(response,gin.H{
// 			"Delivery ID": r.ID,
// 			"Delivery Status": r.DeliveryStatus,
// 		})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"user_id": userID,
// 		"Deliveries": response,
// 	})
// }

// func AddDelivery(c *gin.Context) {
// 	orderIDInterface := c.Param("order_ID")
// 	orderID, err := strconv.Atoi(orderIDInterface)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid order_ID"})
// 	}
// 	var delivery models.Delivery
// 	if err := c.ShouldBindJSON(&delivery); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	delivery.OrderID = uint(orderID)
// 	err = DeliveryTable.Create(c, delivery)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Delivery added",
// 	})
// }
// ///////
// // My project is online_electronic store. For the implementation, the database connection is established in config/db.go file and returns the object through which the connection has been established. Later the instantiation of the DeliveryTable object in controllers/delivery_controller.go file has been done by calling the function in the tables/deliveries_table.go file. The task is to test the working of the controllers/delivery_controller.go using the mockgen. The methods present in the delivery_controller.go