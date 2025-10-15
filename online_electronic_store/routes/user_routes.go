package routes

import (
	"online_electronic_store/handlers"
	"online_electronic_store/interfaces"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	var userTable interfaces.Users = handlers.NewUserTable(db)
	r.POST("/register", userTable.Register)
	r.POST("/login", userTable.Login)
}
