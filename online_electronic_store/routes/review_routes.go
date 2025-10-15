package routes

import (
	"online_electronic_store/handlers"
	"online_electronic_store/interfaces"
	"online_electronic_store/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReviewRoutes(r *gin.Engine, db *gorm.DB) {
	var reviewTable interfaces.Reviews = handlers.NewReviewTable(db)
	r.GET("/reviews/product/:product_id", middleware.AuthMiddleware(), reviewTable.GetReviewOnProductID) 
	r.GET("/reviews", middleware.AuthMiddleware(), reviewTable.GetReviewsOnUserID)
	r.POST("/review/:product_id/create", middleware.AuthMiddleware(), reviewTable.CreateReview)
}
