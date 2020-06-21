package domain

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Todos []Todo
