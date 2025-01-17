package main

import (
	"todo-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //Here router is our server

	//ROUTES
	router.GET("/todos", controller.GetTodos)
	router.GET("/todos/:id", controller.GetTodo)
	router.PATCH("/todos/:id", controller.ToggleTodoStatus)
	router.POST("/todos", controller.AddTodo)

	router.Run("localhost:9090")
}
