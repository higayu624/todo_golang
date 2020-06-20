package database

import (
	"time"

	"github.com/hirokikondo86/clean-architecture-sample/src/app/domain"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindAll() (todos domain.Todos, err error) {
	rows, err := repo.Query("SELECT id, task, limit_date, status FROM todos")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var task string
		var limitDate time.Time
		var status bool
		if err := rows.Scan(&id, &task, &limitDate, &status); err != nil {
			continue
		}

		todos := []domain.Todo{}
		todo := domain.Todo{
			ID:        id,
			Task:      task,
			LimitDate: limitDate,
			Status:    status,
		}
		todos = append(todos, todo)
	}
	return
}

func (repo *TodoRepository) FindById(identifier int) (todo domain.Todo, err error) {
	row, err := repo.Query("SELECT id, task, limit_date, status WHERE id = ? FROM todos", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var task string
	var limitDate time.Time
	var status bool
	row.Next()
	if err = row.Scan(&id, &task, &limitDate, &status); err != nil {
		return
	}
	todo.ID = id
	todo.Task = task
	todo.LimitDate = limitDate
	todo.Status = status
	return
}

func (repo *TodoRepository) Store(todo domain.Todo) (err error) {
	_, err = repo.Execute(
		"INSERT INTO todos (task,limitDate,status) VALUES (?, ?, ?)", todo.Task, todo.LimitDate, todo.Status,
	)
	if err != nil {
		return
	}
	return
}

func (repo *TodoRepository) Update(todo domain.Todo) (err error) {
	_, err = repo.Execute(
		"UPDATE todos SET task = ?, limitDate = ?, status = ? WHERE id = ?", todo.Task, todo.LimitDate, todo.Status, todo.ID,
	)
	if err != nil {
		return
	}
	return
}
