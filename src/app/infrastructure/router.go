package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/higayu624/todo_golang/src/app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	r := gin.Default()

	r.Use(cors.Default())

	todoController := controllers.NewTodoController(NewSqlHandler())

	r.GET("/api/v1/todos", func(c *gin.Context) { todoController.Index(c) })
	r.GET("/api/v1/todos/:id", func(c *gin.Context) { todoController.Show(c) })
	r.POST("/api/v1/todos", func(c *gin.Context) { todoController.Create(c) })
	r.PUT("/api/v1/todos/:id", func(c *gin.Context) { todoController.Update(c) })
	r.DELETE("/api/v1/todos/:id", func(c *gin.Context) { todoController.Destroy(c) })

	Router = r
}
