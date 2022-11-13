package controllers

import (
	"fmt"
	"strconv"

	"github.com/higayu624/todo_golang/src/app/domain"
	"github.com/higayu624/todo_golang/src/app/interfaces/database"
	"github.com/higayu624/todo_golang/src/app/usecase"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) Index(c Context) {
	todos, err := controller.Interactor.Todos()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, todos)
}

func (controller *TodoController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := controller.Interactor.TodoById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, todo)
}

func (controller *TodoController) Create(c Context) {
	var t domain.Todo
	c.ShouldBind(&t)
	todo, err := controller.Interactor.Add(t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, todo)
}

func (controller *TodoController) Update(c Context) {
	var t domain.Todo

	t.ID, _ = strconv.Atoi(c.Param("id"))
	c.ShouldBind(&t)

	todo, err := controller.Interactor.Edit(t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, todo)
}

func (controller *TodoController) Destroy(c Context) {
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.Interactor.Delete(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	res := Response{
		Status:  "success",
		Message: fmt.Sprintf("succeeded in deleting id %d", id),
	}
	c.JSON(200, res)
}
