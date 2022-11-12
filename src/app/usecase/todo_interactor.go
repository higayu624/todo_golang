package usecase
//アプリケーション固有のビジネスルール

import "github.com/higayu624/todo_golang/src/app/domain"

type TodoInteractor struct {
	TodoRepository TodoRepository
	//sqlhandler.goのTodoRepositoryインターフェースを継承している
}

//インターフェースを継承したメソッドの宣言
func (interactor *TodoInteractor) Todos() (todo domain.Todos, err error) {
	todo, err = interactor.TodoRepository.FindAll()
	return
}

func (interactor *TodoInteractor) TodoById(identifier int) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.FindById(identifier)
	return
}

func (interactor *TodoInteractor) Add(t domain.Todo) (todo domain.Todo, err error) {
	identifier, err := interactor.TodoRepository.Store(t)
	if err != nil {
		return
	}
	todo, err = interactor.TodoRepository.FindById(identifier)
	return
}

func (interactor *TodoInteractor) Edit(t domain.Todo) (todo domain.Todo, err error) {
	identifier, err := interactor.TodoRepository.Update(t)
	if err != nil {
		return
	}
	todo, err = interactor.TodoRepository.FindById(identifier)
	return
}

func (interactor *TodoInteractor) Delete(id int) (err error) {
	err = interactor.TodoRepository.Delete(id)
	if err != nil {
		return
	}
	return
}
