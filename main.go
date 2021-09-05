package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Level int16  `json:"level"`
	Daily bool   `json:"daily"`
}

var tasks = []task{
	{ID: 1, Name: "Limpar chão da cozinha", Level: 2, Daily: false},
	{ID: 2, Name: "Limpar chão da varanda", Level: 2, Daily: false},
	{ID: 3, Name: "Limpar caixa de areia", Level: 1, Daily: true},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
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

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.POST("/tasks", postTasks)

	router.Run("localhost:8080")
}
