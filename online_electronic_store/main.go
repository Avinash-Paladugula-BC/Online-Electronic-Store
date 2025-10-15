package main

import (
	"fmt"
	"log"
	"online_electronic_store/config"
	"online_electronic_store/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	db := config.ConnectDB()
	config.DB = config.ConnectDB()

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
