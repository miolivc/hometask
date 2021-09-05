package main

import (
	controllers "com.github.miolivc/hometask/controllers"
	"com.github.miolivc/hometask/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := models.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTaskById)
	router.POST("/tasks", controllers.PostTask)
	router.DELETE("/tasks/:id", controllers.DeleteTaskById)

	router.Run("localhost:8080")
}
