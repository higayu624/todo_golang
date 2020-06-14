package domain

import "time"

type Todo struct {
	ID        int       `json: "id"`
	Task      string    `json: "task"`
	LimitDate time.Time `json: "limitDate"`
	Status    bool      `json: "status"`
}

type Todos []Todo
