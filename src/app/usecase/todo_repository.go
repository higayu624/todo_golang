package usecase

import "github.com/higayu624/todo_golang/src/app/domain"

type TodoRepository interface {
	FindAll() (domain.Todos, error)
	FindById(int) (domain.Todo, error)
	Store(domain.Todo) (int, error)
	Update(domain.Todo) (int, error)
	Delete(int) error
}
