package handlers

import (
	"fmt"
	"log"
	"net/http"
	"online_electronic_store/auth"
	"online_electronic_store/interfaces"
	"online_electronic_store/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserTable struct {
	dB *gorm.DB
}

func NewUserTable(DB *gorm.DB) interfaces.Users {
	return &UserTable{
		dB: DB,
	}
}

func (u *UserTable) Login(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := u.dB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password...."})
		return
	}
	token, err := auth.GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserTable) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// encrypting the password, cost factor of 12 is better to avoid over computation
	encrypted_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		log.Print("Error while generating password", err.Error())
	}
	user.Password = string(encrypted_password)
	// Adding the new user
	fmt.Println("in the user......................", u.dB)
	if err := u.dB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to create user", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}
