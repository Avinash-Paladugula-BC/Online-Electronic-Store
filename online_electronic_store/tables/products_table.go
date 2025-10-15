package tables

// import (
// 	// "net/http"
// 	"online_electronic_store/interfaces"
// 	"online_electronic_store/models"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type ProductTable struct {
// 	dB *gorm.DB
// }


// // To make sure all methods are implemented
// // var _ interfaces.Online_Electronic_Store_DB[models.Product] = &ProductTable{}

// func NewProductTable(DB *gorm.DB) interfaces.Online_Electronic_Store_DB[models.Product] {
// 	return &ProductTable{
// 		dB: DB,
// 	}
// }

// func (r *ProductTable) Create(c *gin.Context, product models.Product) error {
// 	return r.dB.Create(&product).Error
// }

// func (r *ProductTable) GetByID(c *gin.Context, productID uint) (models.Product, error) {
// 	userIDInterface, _ := c.Get("userID")
// 	userID := userIDInterface.(uint)
// 	var product  models.Product
	
// 	result := r.dB.Where("user_id = ? and product_id = ?", userID, productID ).First(&product)
// 	if result.Error != nil{
// 		return product, result.Error
// 	}
// 	return product, nil
// }

// func (r *ProductTable) GetAll(c *gin.Context ) ([]models.Product, error) {
// 	////////
// 	var products []models.Product
// 	result := r.dB.Find(&products)
// 	return products, result.Error
// }

// // func (r *ProductTable) Update(c *gin.Context, data models.Product) error {
	
// // }
