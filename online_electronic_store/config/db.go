package config

import (
	"fmt"
	"log"
	"os"

	"online_electronic_store/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDB() *gorm.DB {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found or couldn't be loaded")
	}
	var DB *gorm.DB
	fmt.Println(os.Getenv("HOST"), "++++++++++++++++++++++++++++,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,")
	creds := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("PORT"))
	DB, err = gorm.Open(postgres.Open(creds), &gorm.Config{})
	if err != nil {
		log.Printf("Error occured : ", err.Error())
	}
	log.Println("Database connection established")

	err = DB.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{}, &models.Payment{}, &models.Review{}, &models.Delivery{})
	if err != nil {
		log.Print("Error while creating tables : ", err.Error())
	}
	log.Print("Database connection established and tables automigrated")
	return DB
}
