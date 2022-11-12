package main

import (
	"fmt"

	"github.com/higayu624/todo_golang/src/app/infrastructure"
)

func main() {
	fmt.Println("sever start")
	infrastructure.Router.Run()
}
