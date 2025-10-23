package interfaces

import (
	"github.com/gin-gonic/gin"
)
type Deliveries interface{
	GetDeliveryDetailsByOrderID(*gin.Context) 
	GetDeliveryDetailsList(*gin.Context)
	AddDelivery(*gin.Context)
}

type Orders interface{
	GetOrders(*gin.Context)
	MakeOrder(*gin.Context)
	UpdateOrder(*gin.Context)
}

type Payments interface{
	GetPaymentsOfUser(*gin.Context)
	CreatePayment(*gin.Context)
	UpdatePayment(*gin.Context)
}

type Products interface{
	GetProducts(*gin.Context)
	GetProductById(*gin.Context)
	GetProductByCategory(*gin.Context)
	CreateProduct(*gin.Context)
	UpdateProductOnId(*gin.Context)
}

type Reviews interface{
	GetReviewOnProductID(*gin.Context)
	GetReviewsOnUserID(*gin.Context)
	CreateReview(*gin.Context)
}

type Users interface{
	Register(*gin.Context)
	Login(*gin.Context)
}