package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

	// Register the IPLogger middleware
    r.Use(middlewares.IPLogger())

    r.GET("/category", controllers.FindCategories)
    r.POST("/category", controllers.CreateCategory)
    r.GET("/category/:id", controllers.FindCategory)
    r.PUT("/category", controllers.UpdateCategory)
    r.DELETE("/category/:id", controllers.DeleteCategory)

    r.GET("/activity", controllers.FindActivities)
    r.POST("/activity", controllers.CreateActivity)
    r.GET("/activity/:id", controllers.FindActivity)
    r.DELETE("/activity/:id", controllers.DeleteActivity)
    
    r.GET("/submission/pending", controllers.FindPendingSubmissions)
    r.GET("/submission/approved", controllers.FindApprovedSubmissions)
    r.GET("/submission/:id", controllers.FindPendingSubmission)
    r.POST("/submission", controllers.CreateSubmission)
    r.PUT("/submission/:id", controllers.UpdateSubmission)
    r.PUT("/submission/inactive/:id", controllers.SetSubmissionInactive)
    r.PUT("/submission/:id/:status", controllers.UpdateSubmissionStatus)


    return r
}
