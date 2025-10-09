package main

import (
	"fmt"
	"log"
	"online-electronic-store/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// this file is just to check if any errors occur while automigration
	DB_NAME := "oes"
	HOST := "localhost"
	PORT := "5432"
	USER := "postgres"
	PASSWORD := "Avi@2004"
	creds := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		HOST, USER, PASSWORD, DB_NAME, PORT)
	var err error
	DB, err = gorm.Open(postgres.Open(creds), &gorm.Config{})
	if err != nil {
		log.Printf("Error occured : ", err.Error())
	}
	log.Println("Database connection established")

	err = DB.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{}, &models.Payment{}, &models.Review{}, &models.Delivery{})
	if err != nil {
		log.Print("Error while creating tables : ", err.Error())
	}

	log.Print("Database connectin established and tables automigrated")

}
