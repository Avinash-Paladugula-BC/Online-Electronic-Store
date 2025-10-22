package config

import (
	"fmt"
	"log"
	"os"

	"online_electronic_store/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// export DB_NAME="oes"
// export HOST="localhost"
// export PORT="5432"
// export USER="postgres"
// export PASSWORD="Avi@2004"


func ConnectDB() *gorm.DB{ // return DB here
	// host := "localhost"
	// port := "5432"
	// user := "postgres"
	// password := "Avi@2004"
	// dbname := "online_electronic_store"
	var DB *gorm.DB//////////////////
	creds := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("PORT"))
	// DB, err := gorm.Open("postgres", creds)
	var err error
	DB, err = gorm.Open(postgres.Open(creds), &gorm.Config{})
	if err != nil {
		log.Printf("Error occured : ", err.Error())
	}
	log.Println("Database connection established")

	// err = DB.AutoMigrate(&models.Product{}, &models.Rating{}, &models.User{})
	err = DB.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{}, &models.Payment{}, &models.Review{}, &models.Delivery{})
	// err = DB.AutoMigrate(&models.Product{}, &models.Review{})
	if err != nil{
		log.Print("Error while creating tables : ",err.Error())
	}
	log.Print("Database connection established and tables automigrated")
	return DB
}
