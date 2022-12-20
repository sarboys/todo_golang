package main

import (
	"github.com/sarboys/todo_golang"
	"github.com/sarboys/todo_golang/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error")
	}
}
