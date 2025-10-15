package routes

import (
	"online_electronic_store/handlers"
	"online_electronic_store/interfaces"
	"online_electronic_store/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymentRoutes(r *gin.Engine, db *gorm.DB) {
	var paymentTable interfaces.Payments = handlers.NewPaymentTable(db)
	r.GET("/payments", middleware.AuthMiddleware(), paymentTable.GetPaymentsOfUser)
	r.POST("/payments/create/:order_id", middleware.AuthMiddleware(), paymentTable.CreatePayment)
	r.PUT("/payments/:payment_id", middleware.AuthMiddleware(), paymentTable.UpdatePayment)
}
