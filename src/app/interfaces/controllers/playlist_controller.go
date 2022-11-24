package controllers

import (
	"github.com/higayu624/todo_golang/src/app/interfaces/requests"
	"github.com/higayu624/todo_golang/src/app/usecase"
)

type PlayListController struct {
	Interactor usecase.PlayListInteractor
}

func NewPlayListController(requestHandler requests.RequestHandler){
	return &PlayListController{
		Interactor: usecase.PlayListInteractor{
			PlayListRepository: &requests.PlayListRepository{
				RequestHandler: requestHandler,
			},
		},
	}
}

func (controller *PlayListController) Show(c Context){
	playlist, err := controller.Interactor.GetPlayList()
	panic("い")
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, playlist)
}