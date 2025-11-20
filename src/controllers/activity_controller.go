package controllers

import (
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindActivities(c *gin.Context) {
    var activities []models.Activity
    models.DB.Find(&activities)
    utils.RespondJSON(c, http.StatusOK, activities)
}

func CreateActivity(c* gin.Context) {
    var input models.Activity
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    var existingCategory models.Category
    if err := models.DB.Where("ID = ?", input.CategoryId).First(&existingCategory).Error; err != nil {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid category ID."})
        return
    }

    models.DB.Create(&input)
    utils.RespondJSON(c, http.StatusCreated, input)
}

func FindActivity(c *gin.Context) {
    var activity models.Activity
    if err := models.DB.Where("id = ?", c.Param("id")).First(&activity).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Activity not found!"})
        return
    }
    utils.RespondJSON(c, http.StatusOK, activity)
}

func DeleteActivity(c *gin.Context) {
    var activity models.Activity
    if err := models.DB.Where("id = ?", c.Param("id")).First(&activity).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Activity not found!"})
        return
    }
    models.DB.Delete(&activity)
    utils.RespondJSON(c, http.StatusOK, activity)
}
