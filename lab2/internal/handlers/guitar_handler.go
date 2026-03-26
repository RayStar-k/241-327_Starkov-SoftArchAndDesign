package handlers

import (
	"net/http"
	"strconv"

	"guitarshop/internal/database"
	"guitarshop/internal/models"

	"github.com/gin-gonic/gin"
)

func GetGuitars(c *gin.Context) {
	var guitars []models.Guitar
	query := database.DB

	if brand := c.Query("brand"); brand != "" {
		query = query.Where("brand = ?", brand)
	}
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}
	if inStock := c.Query("in_stock"); inStock != "" {
		query = query.Where("in_stock = ?", inStock)
	}

	result := query.Find(&guitars)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, guitars)
}

func GetGuitar(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var guitar models.Guitar
	result := database.DB.First(&guitar, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
		return
	}

	c.JSON(http.StatusOK, guitar)
}

func CreateGuitar(c *gin.Context) {
	var guitar models.Guitar

	if err := c.ShouldBindJSON(&guitar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&guitar)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, guitar)
}

func UpdateGuitar(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var guitar models.Guitar
	if result := database.DB.First(&guitar, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
		return
	}

	if err := c.ShouldBindJSON(&guitar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&guitar)

	c.JSON(http.StatusOK, guitar)
}

func DeleteGuitar(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result := database.DB.Delete(&models.Guitar{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Guitar deleted successfully"})
}
