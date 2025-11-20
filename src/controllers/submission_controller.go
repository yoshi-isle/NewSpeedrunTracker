package controllers

import (
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPendingSubmissions(c *gin.Context) {
    var submissions []models.Submission
    models.DB.Where("is_approved = ? AND is_active = ?", false, true).Find(&submissions)
    utils.RespondJSON(c, http.StatusOK, submissions)
}

func FindApprovedSubmissions(c *gin.Context) {
    var submissions []models.Submission
    models.DB.Where("is_approved = ? AND is_active = ?", true, true).Find(&submissions)
    utils.RespondJSON(c, http.StatusOK, submissions)
}

func FindPendingSubmission(c *gin.Context) {
    var submission models.Submission
    if err := models.DB.Where("id = ?", c.Param("id")).First(&submission).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Submission not found!"})
        return
    }
    utils.RespondJSON(c, http.StatusOK, submission)
}

func CreateSubmission(c* gin.Context) {
    var input models.Submission
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    var existingActivity models.Category
    if err := models.DB.Where("ID = ?", input.ActivityId).First(&existingActivity).Error; err != nil {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": "Invalid activity ID."})
        return
    }

    input.IsApproved = false
    input.IsActive = true
    models.DB.Create(&input)
    utils.RespondJSON(c, http.StatusCreated, input)
}

func UpdateSubmission(c* gin.Context) {
    var submission models.Submission
    if err := models.DB.Where("id = ?", c.Param("id")).First(&submission).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Submission not found!"})
        return
    }

    var input models.Submission
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondJSON(c, http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Model(&submission).Updates(input)
    utils.RespondJSON(c, http.StatusOK, submission)
}

func UpdateSubmissionStatus(c* gin.Context) {
    var submission models.Submission
    if err := models.DB.Where("id = ?", c.Param("id")).First(&submission).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Submission not found!"})
        return
    }

    approved := c.Param("status") == "true"
    if !approved {
        models.DB.Model(&submission).Update("is_active", false)
    }

    models.DB.Model(&submission).Update("is_approved", approved)
    utils.RespondJSON(c, http.StatusOK, submission)
}

func SetSubmissionInactive(c* gin.Context) {
    var submission models.Submission
    if err := models.DB.Where("id = ?", c.Param("id")).First(&submission).Error; err != nil {
        utils.RespondJSON(c, http.StatusNotFound, gin.H{"error": "Submission not found!"})
        return
    }
    models.DB.Model(&submission).Update("is_active", false)
    utils.RespondJSON(c, http.StatusOK, submission)
}

