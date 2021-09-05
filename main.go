package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Level     int    `json:"level"`
	Daily     bool   `json:"daily"`
	Completed bool   `json:"completed"`
}

var tasks = []task{
	{ID: 1, Name: "Limpar chão da cozinha", Level: 2, Daily: false, Completed: false},
	{ID: 2, Name: "Limpar chão da varanda", Level: 2, Daily: false, Completed: false},
	{ID: 3, Name: "Limpar caixa de areia", Level: 1, Daily: true, Completed: false},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "task ID must be an int value"})
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func postTasks(c *gin.Context) {
	var newTask task

	// Call BindJSON to bind the received JSON
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	// Add the new task to the slice.
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func deleteTaskById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "task ID must be an int value"})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "task was deleted successfully"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskById)
	router.POST("/tasks", postTasks)
	router.DELETE("/tasks/:id", deleteTaskById)

	router.Run("localhost:8080")
}
