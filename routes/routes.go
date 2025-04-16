package routes

import (
	"github.com/gin-gonic/gin"
	"project/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	taskRoutes := r.Group("/api")
	{
		taskRoutes.GET("/mood", controllers.GetMood)
		taskRoutes.POST("/update_mood", controllers.UpdateMood)
	}
}