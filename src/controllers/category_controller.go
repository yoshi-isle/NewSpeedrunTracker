package controllers

import (
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindCategories(c *gin.Context) {
    var categories []models.Category
    models.DB.Find(&categories)
    utils.RespondJSON(c, http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {

    var input models.Category
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if input.Name == "" {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Name cannot be empty."})
        return
    }

    var existingCategory models.Category
    if err := models.DB.Where("name = ?", input.Name).First(&existingCategory).Error; err == nil {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Category name already exists."})
        return
    }
    category := models.Category{Name: input.Name}
    models.DB.Create(&category)
    utils.RespondJSON(c, http.StatusCreated, category)
}

func FindCategory(c *gin.Context) {
    var category models.Category
    if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Category not found!"})
        return
    }
    utils.RespondJSON(c, http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
    var category models.Category
    if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Category not found!"})
        return
    }

    var input models.Category
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Model(&category).Updates(input)
    utils.RespondJSON(c, http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
    var category models.Category
    if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Category not found!"})
        return
    }
    models.DB.Delete(&category)
    utils.RespondJSON(c, http.StatusOK, category)
}
