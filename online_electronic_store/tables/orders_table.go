package tables

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"online_electronic_store/interfaces"
// 	"online_electronic_store/models"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type OrderTable struct {
// 	dB *gorm.DB
// }

// // To make sure all methods are implemented
// // var _ interfaces.Online_Electronic_Store_DB[models.Order] = &OrderTable{}

// func NewOrderTable(DB *gorm.DB) interfaces.Online_Electronic_Store_DB[models.Order] {
// 	return &OrderTable{
// 		dB: DB,
// 	}
// }

// func (r *OrderTable) Create(c *gin.Context, order models.Order) error {
// 	return r.dB.Create(&order).Error
// }

// func (r *OrderTable) GetByID(c *gin.Context, id uint) (models.Order, error) {
// 	// unimlemented
// 	fmt.Println("Not implemented")
// 	var order models.Order
// 	return order, errors.New("method not implemented")
// }

// func (r *OrderTable) GetAll(c *gin.Context) ([]models.Order, error) {
// 	var orders []models.Order
// 	userIDInterface, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
// 		return orders, nil
// 	}
// 	userID := userIDInterface.(uint)
// 	err := r.dB.Where("user_id = ?", userID).Find(&orders).Error
// 	return orders, err
// }

// // func (r *OrderTable) Update(order models.Order, order_id uint) error {

// // }
