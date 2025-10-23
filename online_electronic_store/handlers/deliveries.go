package handlers

import (
	"net/http"
	"online_electronic_store/interfaces"
	"online_electronic_store/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type deliveryTable struct {
	dB *gorm.DB
}

func NewDeliveryTable(DB *gorm.DB) interfaces.Deliveries {
	return &deliveryTable{
		dB: DB,
	}
}


func (d *deliveryTable) OrderDeliveryDetailsFromDB(orderID uint)(models.Delivery, error){
	var delivery models.Delivery
	err := d.dB.Where("order_id = ?", orderID).First(&delivery).Error
	return delivery,err
}

func (d *deliveryTable) GetDeliveryDetailsByOrderID(c *gin.Context) {
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	orderIDInterface := c.Param("order_id")
	order_id, _ := strconv.Atoi(orderIDInterface)

	var delivery models.Delivery
	delivery, err := d.OrderDeliveryDetailsFromDB(uint(order_id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch delivery details"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Delivery ID":     delivery.ID,
		"Delivery status": delivery.DeliveryStatus,
	})
}


func (d *deliveryTable) GetAllDeliveriesOfUser(userID uint) ([]models.Delivery, error){
	var deliveries []models.Delivery
	err := d.dB.
		Joins("JOIN orders ON orders.id = deliveries.order_id").
		Where("orders.user_id = ?", userID).
		Find(&deliveries).Error
	return deliveries, err
}
func (d *deliveryTable) GetDeliveryDetailsList(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	userID := userIDInterface.(uint)
	var deliveries []models.Delivery

	deliveries, err := d.GetAllDeliveriesOfUser(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch delivery details"})
		return
	}
	response := []gin.H{}
	for _, r := range deliveries{
		response=append(response,gin.H{
			"Delivery ID": r.ID,
			"Delivery Status": r.DeliveryStatus,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"Deliveries": response,
	})
}

func (d *deliveryTable) InsertDeliveryIntoDB(delivery models.Delivery) error{
	return d.dB.Create(&delivery).Error
}
func (d *deliveryTable) AddDelivery(c *gin.Context) {
	orderIDInterface := c.Param("order_ID")
	orderID, err := strconv.Atoi(orderIDInterface)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid order_ID"})
	}
	var delivery models.Delivery
	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	delivery.OrderID = uint(orderID)

	err = d.InsertDeliveryIntoDB(delivery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Delivery added",
	})
}

