package main

import (
	"log"
	todo_app "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run("1000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err)
	}
}
