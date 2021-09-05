package controllers

import (
	"net/http"
	"strconv"

	"com.github.miolivc/hometask/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "task ID must be an int value"})
		return
	}

	var task models.Task
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, task)
}

func PostTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error parsing JSON"})
		return
	}

	if err := db.Create(&task); err.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error})
		return
	}

	c.IndentedJSON(http.StatusCreated, task)
}

func DeleteTaskById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "task ID must be an int value"})
		return
	}

	var task models.Task
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	db.Delete(&task)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}
