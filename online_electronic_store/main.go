package main

import (
	"fmt"
	"log"
	"online_electronic_store/config"
	"online_electronic_store/routes"

	"github.com/gin-gonic/gin"
)

// func init(){

// }

func main() {
	db := config.ConnectDB()
	config.DB = config.ConnectDB()
	// create all

	// controller.Delivery_table = tables.NewDeliveryTable(config.DB)
	// controller.InstantiateDeliveryTable(db)
	// controller.InstantiateOrderTable(db)
	// controller.InstantiatePaymentTable(db)
	// controller.InstantiateReviewTable(db)
	// controller.InstantiateUserTable()
	// controller.InstantiateProductTable(db)

	r := gin.Default()
	routes.ProductRoutes(r, db)
	routes.UserRoutes(r, db)
	routes.OrderRoutes(r, db)
	routes.PaymentRoutes(r, db)
	routes.ReviewRoutes(r, db)
	routes.DeliveryRoutes(r, db)

	port := "8080"

	fmt.Printf("Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error occured while starting the server: ", err)
	}

}
