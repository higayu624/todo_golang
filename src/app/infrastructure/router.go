package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/higayu624/todo_golang/src/app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
 
	todoController := controllers.NewTodoController(NewSqlHandler())
	playlistController := controllers.NewPlayListController()

	router.GET("/api/v1/todos", func(c *gin.Context) { todoController.Index(c) })
	router.GET("/api/v1/todos/:id", func(c *gin.Context) { todoController.Show(c) })
	router.POST("/api/v1/todos", func(c *gin.Context) { todoController.Create(c) })
	router.PUT("/api/v1/todos/:id", func(c *gin.Context) { todoController.Update(c) })
	router.DELETE("/api/v1/todos/:id", func(c *gin.Context) { todoController.Destroy(c) })
	router.GET("/api/playlists", playlistController.Show(c))

	Router = r
}
