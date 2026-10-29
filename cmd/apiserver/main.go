package main

import (
	"fmt"

	"github.com/limoore21/go-monolith/internal/repository"
	"github.com/limoore21/go-monolith/internal/service"
)

func main() {
	repo := repository.User{Firstname: "Maxim", Surname: "Burnov"}
	service.NewUserService(repo)
	fmt.Println(repo)
}
