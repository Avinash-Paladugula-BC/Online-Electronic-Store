package tables

// import (
// 	"online_electronic_store/interfaces"
// 	"online_electronic_store/models"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type ReviewTable struct {
// 	dB *gorm.DB
// }

// func NewReviewTable(DB *gorm.DB) interfaces.Online_Electronic_Store_DB[models.Review] {
// 	return &ReviewTable{
// 		dB: DB,
// 	}
// }

// // To make sure all methods are implemented
// // var _ interfaces.Online_Electronic_Store_DB[models.Review] = &ReviewTable{}

// // func NewReviewTable(DB *gorm.DB) interfaces.Online_Electronic_Store_DB[models.Review] {
// // 	return &ReviewTable{
// // 		dB: DB,
// // 	}
// // }

// func (r *ReviewTable) Create(c *gin.Context, review models.Review) error {
// 	return r.dB.Create(&review).Error
// }

// func (r *ReviewTable) GetByID(c *gin.Context, productID uint) (models.Review, error) {
// 	userIDInterface, _ := c.Get("userID")
// 	userID := userIDInterface.(uint)
// 	productIDInterface, _ := c.Get("productID")
// 	productID, _ = productIDInterface.(uint)
// 	var review models.Review

// 	err := r.dB.Where("user_id = ? and product_id = ?",userID, productID).Find(&review)
// 	return review, err.Error
// }

// func (r *ReviewTable) GetAll(c *gin.Context ) ([]models.Review, error) {
// 	//get all the review given by the user
// 	userIDInterface, _ := c.Get("userID")
// 	userID := userIDInterface.(uint)
// 	var reviews []models.Review
// 	err := r.dB.Where("user_id = ?", userID).Find(&reviews).Error
// 	return reviews, err;

// }

