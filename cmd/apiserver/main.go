package main

import (
	"fmt"
	"net/http"

	"github.com/limoore21/go-monolith/internal/handler"
	"github.com/limoore21/go-monolith/internal/repository"
	"github.com/limoore21/go-monolith/internal/service"
)

func main() {
	repo := repository.NewGuestRepository()
	swr := service.NewUserService(repo)

	userHandler := handler.NewUserHandler(swr)

	http.HandleFunc("/user", userHandler.GetUser)

	fmt.Println("Сервер запущен на порту :8082...")
	http.ListenAndServe(":8082", nil)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Printf("Не удалось запустить сервер: %v\n", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Корень сервера работает!")
	})
}
