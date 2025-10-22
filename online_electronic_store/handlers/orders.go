package handlers

import (
	"net/http"
	// "online_electronic_store/config"
	"online_electronic_store/interfaces"
	"online_electronic_store/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type orderTable struct {
	dB *gorm.DB
}

func NewOrderTable(DB *gorm.DB) interfaces.Orders {
	return &orderTable{
		dB: DB,
	}
}

func (o *orderTable) GetOrders(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	userID := userIDInterface.(uint)
	var orders []models.Order
	err := o.dB.Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"orders":  orders,
	})
}
func (o *orderTable) MakeOrder(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	userID := userIDInterface.(uint)
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.UserID = userID
	err := o.dB.Create(&order).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
	})
}
func (o *orderTable) UpdateOrder(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	userID := userIDInterface.(uint)
	orderIDParam := c.Param("order_id")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}
	var order models.Order
	if err := o.dB.First(&order, "ID = ? AND user_id = ?", orderID, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
		return
	}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// order.UserID = userID
	if err := o.dB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order updated successfully",
		"order":   order,
	})
}
