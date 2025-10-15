package tables

// import (
// 	"online_electronic_store/interfaces"
// 	"online_electronic_store/models"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type DeliveryTable struct {
// 	dB *gorm.DB
// }

// func NewDeliveryTable(DB *gorm.DB) interfaces.Online_Electronic_Store_DB[models.Delivery] {
// 	return &DeliveryTable{
// 		dB: DB,
// 	}
// }

// // // To make sure all methods are implemented
// // var _ interfaces.Online_Electronic_Store_DB[models.Delivery] = &DeliveryTable{}

// func (r *DeliveryTable) Create(c *gin.Context, delivery models.Delivery) error {
// 	return r.dB.Create(&delivery).Error
// }
// func (r *DeliveryTable) GetByID(c *gin.Context, delivery_id uint) (models.Delivery, error) {
// 	// userIDInterface, _ := c.Get("userID")
// 	// userID := userIDInterface.(uint)
// 	var delivery models.Delivery
// 	r.dB.Where("ID = ?", delivery_id).First(&delivery)
// 	return delivery, nil
// }

// // query in gorm: get all the records from the table deliveries which are ordered by a particular user where the user id based on the user id. The delivery table contains the order id. In the orders table we have user id. So only when the user id in the order table matches with the
// // get all records from deliveries table. deliveries table contains the order_id. orders table contains the user_id. when the user_id in the orders table matches with the passed user_id then retrieve all the records from the deliveries table
// type DeliveryWithAddress struct {
// 	models.Delivery
// 	ShippingAddress string
// }

// func (r *DeliveryTable) GetAll(c *gin.Context) ([]models.Delivery, error) {
// 	userIDInterface, _ := c.Get("userID")
// 	userID := userIDInterface.(uint)
// 	var deliveries []models.Delivery
// 	// err := r.dB.Where("user_id = ?", userID).Find(&deliveries).Error
// 	err := r.dB.
// 		Joins("JOIN orders ON orders.id = deliveries.order_id").
// 		Where("orders.user_id = ?", userID).
// 		Find(&deliveries).Error
// 	return deliveries, err
// 	// var deliveries []DeliveryWithAddress
// 	// err := r.dB.
// 	// 	Table("deliveries").
// 	// 	Select("deliveries.*, orders.shipping_address").
// 	// 	Joins("JOIN orders ON orders.id = deliveries.order_id").
// 	// 	Where("orders.user_id = ?", userID).
// 	// 	Scan(&deliveries).Error
// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"delivery with address": deliveries,
// 	// })
// 	// var deliverie []models.Delivery
// 	// return deliverie, err
// }
