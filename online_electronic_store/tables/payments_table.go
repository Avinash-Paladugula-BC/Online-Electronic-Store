package tables

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"online_electronic_store/interfaces"
// 	"online_electronic_store/models"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type PaymentTable struct {
// 	dB *gorm.DB
// }

// // To make sure all methods are implemented
// // var _ interfaces.Online_Electronic_Store_DB[models.Payment] = &PaymentTable{}

// func NewPaymentTable(DB *gorm.DB) interfaces.Online_Electronic_Store_DB[models.Payment] {
// 	return &PaymentTable{
// 		dB: DB,
// 	}
// }

// func (r *PaymentTable) Create(c *gin.Context, payment models.Payment) error {
// 	order_ID_string := c.Param("order_ID")
// 	orderID, err := strconv.Atoi(order_ID_string)
// 	payment.OrderID = uint(orderID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// 	err = r.dB.Create(&payment).Error
// 	return err
// }

// func (r *PaymentTable) GetByID(c *gin.Context, id uint) (models.Payment, error) {
// 	fmt.Println("Not implemented")
// 	var payment models.Payment
// 	return payment, errors.New("not implemented")
// }

// func (r *PaymentTable) GetAll(c *gin.Context) ([]models.Payment, error) {
// 	var payments []models.Payment
// 	userIDInterface, _ := c.Get("userID")
// 	userID := userIDInterface.(uint)

// 	err := r.dB.Where("user_id = ?", userID).Find(&payments).Error
// 	if err != nil {
// 		return payments, err
// 	}
// 	return payments, nil

// }

// // func (r *PaymentTable) Update(c *gin.Context, data models.Payment) error {

// // }
