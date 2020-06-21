package database

import (
	"github.com/hirokikondo86/clean-architecture-sample/src/app/domain"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindAll() (todos domain.Todos, err error) {

	rows, err := repo.Query("SELECT id, title FROM todos")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var id int
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			continue
		}
		todo := domain.Todo{
			ID:    id,
			Title: title,
		}
		todos = append(todos, todo)
	}
	return
}

func (repo *TodoRepository) FindById(identifier int) (todo domain.Todo, err error) {
	row, err := repo.Query("SELECT id, title FROM todos WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var title string
	row.Next()
	if err = row.Scan(&id, &title); err != nil {
		return
	}
	todo.ID = id
	todo.Title = title
	return
}

func (repo *TodoRepository) Store(t domain.Todo) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO todos (title) VALUES (?)", t.Title,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *TodoRepository) Update(t domain.Todo) (id int, err error) {
	_, err = repo.Execute(
		"UPDATE todos SET title = ? WHERE id = ?", t.Title, t.ID,
	)
	if err != nil {
		return
	}
	id = int(t.ID)
	return
}

func (repo *TodoRepository) Delete(identifier int) (err error) {
	_, err = repo.Execute(
		"DELETE FROM todos WHERE id = ?", identifier,
	)
	if err != nil {
		return
	}
	return
}
