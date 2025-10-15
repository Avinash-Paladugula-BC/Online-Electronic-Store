package controller

// import (
// 	"fmt"
// 	"net/http"
// 	"online_electronic_store/config"
// 	"online_electronic_store/interfaces"
// 	"online_electronic_store/models"
// 	"online_electronic_store/tables"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// var OrderTable interfaces.Online_Electronic_Store_DB[models.Order]

// func InstantiateOrderTable(dbb *gorm.DB) {
// 	OrderTable = tables.NewOrderTable(dbb)
// 	fmt.Println("Instantiated order table", OrderTable)
// }

// func GetOrders(c *gin.Context) {
// 	userIDInterface, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
// 		return
// 	}
// 	userID := userIDInterface.(uint)
// 	var orders []models.Order
// 	orders, err := OrderTable.GetAll(c)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"user_id": userID,
// 		"orders":  orders,
// 	})
// }

// func MakeOrder(c *gin.Context) {
// 	userIDInterface, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
// 		return
// 	}
// 	userID := userIDInterface.(uint)
// 	var order models.Order
// 	if err := c.ShouldBindJSON(&order); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	order.UserID = userID
// 	err := OrderTable.Create(c, order)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Order created successfully",
// 	})
// }

// func UpdateOrder(c *gin.Context) {
// 	// this is a PUT method. order_id is the parameter for which the order needs to be updated.
// 	userIDInterface, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
// 		return
// 	}
// 	userID := userIDInterface.(uint)
// 	orderIDParam := c.Param("order_id")
// 	orderID, err := strconv.Atoi(orderIDParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
// 		return
// 	}
// 	var order models.Order
// 	if err := config.DB.First(&order, "ID = ? AND user_id = ?", orderID, userID).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
// 		return
// 	}
// 	if err := c.ShouldBindJSON(&order); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	order.UserID = userID
// 	if err := config.DB.Save(&order).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Order updated successfully",
// 		"order":   order,
// 	})
// }
