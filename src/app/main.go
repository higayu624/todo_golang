package main

import (
	"fmt"

	"github.com/hirokikondo86/clean-architecture-sample/src/app/infrastructure"
)

func main() {
	fmt.Println("sever start")
	infrastructure.Router.Run()
}
