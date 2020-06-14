package controllers

import (
	"strconv"

	"github.com/hirokikondo86/clean-architecture-sample/src/app/domain"
	"github.com/hirokikondo86/clean-architecture-sample/src/app/interface/database"
	"github.com/hirokikondo86/clean-architecture-sample/src/app/usecase"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: usecase.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) Create(c Context) {
	todo := domain.Todo{}
	c.Bind(&todo)
	err := controller.Interactor.Add(todo)
	if err != nil {
		c.JSON(500, err)
		return
	}
}

func (controller *TodoController) Index(c Context) {
	todos, err := controller.Interactor.Todos()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, todos)
}

func (controller *TodoController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := controller.Interactor.TodoById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, todo)
}
