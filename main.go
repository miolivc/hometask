package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Level int16  `json:"level"`
	Daily bool   `json:"daily"`
}

var tasks = []task{
	{ID: "1", Name: "Limpar chão da cozinha", Level: 2, Daily: false},
	{ID: "1", Name: "Limpar chão da varanda", Level: 2, Daily: false},
	{ID: "1", Name: "Limpar caixa de areia", Level: 1, Daily: true},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)

	router.Run("localhost:8080")
}
