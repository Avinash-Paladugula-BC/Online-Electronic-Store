package controller

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"online_electronic_store/config"
// 	"online_electronic_store/interfaces"
// 	"online_electronic_store/models"
// 	"online_electronic_store/tables"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// var ProductTable interfaces.Online_Electronic_Store_DB[models.Product]

// func InstantiateProductTable(dbb *gorm.DB) {
// 	ProductTable = tables.NewProductTable(dbb)
// 	fmt.Println("Instantiated product table", ProductTable)
// }

// func GetProducts(c *gin.Context) {
// 	// Login not required
// 	var products []models.Product
// 	products, err := ProductTable.GetAll(c)
// 	result := config.DB.Find(&products)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
// 	}
// 	response := []gin.H{}
// 	for _, res := range products {
// 		response = append(response, gin.H{
// 			"product ID":       res.ID,
// 			"Product Name":     res.Name,
// 			"Product Category": res.Category,
// 			"Product Price":    res.Price,
// 			"Product Brand":    res.Brand,
// 			"Quantity":         res.Quantity,
// 		})
// 	}
// 	c.JSON(http.StatusOK, response)
// }

// func GetProductById(c *gin.Context) {
// 	// Login not required
// 	id := c.Param("id")
// 	productID, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// 	var product models.Product
// 	product, err = ProductTable.GetByID(c, uint(productID))
// 	// config.DB.First(&product, id)
// 	if err := config.DB.Preload("Reviews").First(&product, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
// 		return
// 	}
// 	// response := gin.H{}
// 	// response.id=product.ID
// 	// response.laksd=alskd
// 	// for _,
// 	c.JSON(http.StatusOK, product)
// }

// func GetProductByCategory(c *gin.Context) {
// 	category := c.Param("category")
// 	var products []models.Product
// 	// *************** search on the category column not on the PK
// 	if err := config.DB.Where("category = ?", category).Find(&products).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"category": category, "products": products})
// }

// //	if err := config.DB.Where("user_id = ?", userID).Find(&payments).Error; err != nil {
// //		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
// //		return
// //	}
// func CreateProduct(c *gin.Context) {
// 	userIDInterface, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
// 		return
// 	}
// 	userID := userIDInterface.(uint)
// 	var user models.User
// 	if err := config.DB.First(&user, userID).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
// 		return
// 	}
// 	if user.Role != "admin" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can add products"})
// 		return
// 	}

// 	var product models.Product
// 	// product = ProductTable.Create(c,)
// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	err := ProductTable.Create(c, product)
// 	// result :=
// 	if err.Error != nil {
// 		log.Println(err.Error)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while adding the product"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Product created"})
// }

// func UpdateProductOnId(c *gin.Context) {
// 	userIDInterface, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
// 		return
// 	}
// 	userID := userIDInterface.(uint)
// 	// check if admin or not
// 	var user models.User
// 	if err := config.DB.First(&user, userID).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
// 		return
// 	}
// 	if user.Role != "admin" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can add products"})
// 		return
// 	}

// 	productIDParam := c.Param("product_id")
// 	productID, err := strconv.Atoi(productIDParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
// 		return
// 	}
// 	var product models.Product
// 	if err := config.DB.First(&product, "id = ?", productID).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Product record not found"})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Product details"})
// 		return
// 	}
// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	product.ID = uint(productID)
// 	if err := config.DB.Save(&product).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Product updated successfully",
// 		"Product": product,
// 	})
// }
