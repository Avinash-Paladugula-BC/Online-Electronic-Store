package interfaces

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

// type Online_Electronic_Store_DB[T any] interface {
// 	Create(c *gin.Context, data T) error
// 	GetByID(c *gin.Context, id uint) (T, error) // get record on id
// 	GetAll(c *gin.Context) ([]T, error)        // get all records
// 	// Update(data T) error // update on id--need to check if its valid for all tables
// 	// Delete(id int) ()
// }

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