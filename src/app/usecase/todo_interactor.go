package usecase

import (
	"github.com/hirokikondo86/clean-architecture-sample/src/app/domain"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Todos() (todos domain.Todos, err error) {
	todos, err = interactor.TodoRepository.FindAll()
	return
}

func (interactor *TodoInteractor) TodoById(id int) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.FindById(id)
	return
}

func (interactor *TodoInteractor) Add(todo domain.Todo) (err error) {
	err = interactor.TodoRepository.Store(todo)
	return
}

func (interactor *TodoInteractor) Edit(todo domain.Todo) (err error) {
	err = interactor.TodoRepository.Update(todo)
	return
}
