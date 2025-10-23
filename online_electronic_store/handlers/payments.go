package handlers

import (
	"net/http"
	"online_electronic_store/interfaces"
	"online_electronic_store/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type paymentTable struct {
	dB *gorm.DB
}

func NewPaymentTable(DB *gorm.DB) interfaces.Payments {
	return &paymentTable{
		dB: DB,
	}
}

func (p *paymentTable) GetPaymentsOfUser(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	userID := userIDInterface.(uint)
	var payments []models.Payment
	err := p.dB.Where("user_id = ?", userID).Find(&payments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
		return
	}
	response := []gin.H{}
	for _, res := range payments {
		response = append(response, gin.H{
			"Payment ID":     res.ID,
			"Order ID":       res.OrderID,
			"Payment Status": res.PaymentStatus,
			"Payment Method": res.PaymentMethod,
			"Transaction ID": res.TransactionID,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id":  userID,
		"payments": response,
	})
}
func (p *paymentTable) CreatePayment(c *gin.Context) {
	order_ID_string := c.Param("order_ID")
	orderID, err := strconv.Atoi(order_ID_string)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment.OrderID = uint(orderID)
	err = p.dB.Create(&payment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Payment"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Payment added successfully",
	})
}
func (p *paymentTable) UpdatePayment(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	userID := userIDInterface.(uint)
	paymentIDParam := c.Param("payment_id")
	paymentID, err := strconv.Atoi(paymentIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}
	var payment models.Payment
	if err := p.dB.First(&payment, "ID = ? AND user_id = ?", paymentID, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Payment details"})
		return
	}
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment.UserID = userID
	if err := p.dB.Save(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment updated successfully",
		"payment": payment,
	})
}
