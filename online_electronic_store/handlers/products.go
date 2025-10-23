package handlers

import (
	"net/http"
	// "online_electronic_store/config"
	"online_electronic_store/interfaces"
	"online_electronic_store/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productTable struct {
	dB *gorm.DB
}

func NewProductTable(DB *gorm.DB) interfaces.Products {
	return &productTable{
		dB: DB,
	}
}

func (p *productTable) GetProducts(c *gin.Context) {
	var products []models.Product
	result := p.dB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	}
	response := []gin.H{}
	for _, res := range products {
		response = append(response, gin.H{
			"product ID":       res.ID,
			"Product Name":     res.Name,
			"Product Category": res.Category,
			"Product Price":    res.Price,
			"Product Brand":    res.Brand,
			"Quantity":         res.Quantity,
		})
	}
	c.JSON(http.StatusOK, response)
}
func (p *productTable) GetProductById(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var product models.Product

	if err := p.dB.Preload("Reviews").Where("product_id = ?", productID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
func (p *productTable) GetProductByCategory(c *gin.Context) {
	category := c.Param("category")
	var products []models.Product
	if err := p.dB.Where("category = ?", category).Find(&products).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"category": category, "products": products})
}
func (p *productTable) CreateProduct(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	userID := userIDInterface.(uint)
	var user models.User
	if err := p.dB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can add products"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := p.dB.Create(&product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while adding the product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product created"})
}
func (p *productTable) UpdateProductOnId(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	userID := userIDInterface.(uint)
	var user models.User
	if err := p.dB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can add products"})
		return
	}

	productIDParam := c.Param("product_id")
	productID, err := strconv.Atoi(productIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	var product models.Product
	if err := p.dB.First(&product, "id = ?", productID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Product details"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.dB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"Product": product,
	})
}
