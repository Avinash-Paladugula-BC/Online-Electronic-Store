package handlers

import (
	"net/http"
	"online_electronic_store/interfaces"
	"online_electronic_store/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewTable struct {
	dB *gorm.DB
}

func NewReviewTable(DB *gorm.DB) interfaces.Reviews {
	return &ReviewTable{
		dB: DB,
	}
}

func (r *ReviewTable) GetReviewOnProductID(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	userID := userIDInterface.(uint)
	productIDInterface := c.Param("product_id")
	productID, _ := strconv.Atoi(productIDInterface)
	var review models.Review
	// review, err := ReviewTable.GetByID(c, uint(product_id))
	err := r.dB.Where("user_id = ? and product_id = ?", userID, productID).Find(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch review details"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"review":  review,
	})
}
func (r *ReviewTable) GetReviewsOnUserID(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	userID := userIDInterface.(uint)
	var reviews []models.Review
	if err := r.dB.Where("user_id = ?", userID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch reviews"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"reviews": reviews,
	})
}
func (r *ReviewTable) CreateReview(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	userID := userIDInterface.(uint)
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	review.UserID = userID
	err := r.dB.Create(&review).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Review added successfully",
	})
}
