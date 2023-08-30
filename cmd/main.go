package main

import (
	"log"

	"github.com/khlkvmgmd/todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run(":8000"); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
