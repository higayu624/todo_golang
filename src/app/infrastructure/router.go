package infrastructure

import (
	"github.com/hirokikondo86/clean-architecture-sample/src/app/interface/controllers"

	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	todoController := controllers.NewTodoController(NewSqlHandler())

	router.GET("/todos", func(c *gin.Context) { todoController.Index(c) })
	router.GET("/todos/:id", func(c *gin.Context) { todoController.Show(c) })
	router.POST("/todos", func(c *gin.Context) { todoController.Create(c) })
	router.PUT("/todos/:id", func(c *gin.Context) { todoController.Update(c) })

	Router = router
}
