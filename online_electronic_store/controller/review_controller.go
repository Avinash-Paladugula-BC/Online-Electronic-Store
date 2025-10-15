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

// var ReviewTable interfaces.Online_Electronic_Store_DB[models.Review]

// func InstantiateReviewTable(dbb *gorm.DB) {
// 	ReviewTable = tables.NewReviewTable(dbb)
// 	fmt.Println("Instantiated Review table", ReviewTable)
// }

// func GetReviewsOnProductID(c *gin.Context) {
// 	// * this doesnt require the login
// 	// userIDInterface, exists := c.Get("userID")
// 	// if !exists{
// 	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
// 	// 	return
// 	// }
// 	// userID := userIDInterface.(uint)
// 	productID := c.Param("product_id")
// 	var reviews []models.Review
// 	if err := config.DB.Where("product_id=?", productID).Find(&reviews).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "faciled to fetch the reviews"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, reviews)
// }

// func GetReviewOnProductID(c *gin.Context) {
// 	userIDInterface, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
// 		return
// 	}
// 	userID := userIDInterface.(uint)
// 	productIDInterface := c.Param("product_id")
// 	product_id, _ := strconv.Atoi(productIDInterface)
// 	review, err := ReviewTable.GetByID(c, uint(product_id))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch review details"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"user_id": userID,
// 		"review":  review,
// 	})
// }

// func GetReviewsOnUserID(c *gin.Context) {
// 	// purpose: to get the details of the
// 	userIDInterface, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
// 		return
// 	}
// 	userID := userIDInterface.(uint)
// 	var reviews models.Review
// 	if err := config.DB.Where("user_id = ?", userID).Find(&reviews).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch reviews"})
// 		return
// 	}
// 	reviews, err := ReviewTable.GetByID(c, userID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"user_id": userID,
// 		"reviews": reviews,
// 	})
// }

// func CreateReview(c *gin.Context) {
// 	userIDInterface, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
// 		return
// 	}
// 	userID := userIDInterface.(uint)
// 	var review models.Review
// 	if err := c.ShouldBindJSON(&review); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	review.UserID = userID
// 	err := ReviewTable.Create(c, review)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Review added successfully",
// 	})
// }
