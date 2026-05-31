package main

import (
	"net/http"
	"project/handlers"
	"project/logger"
	"project/storage"
)

func main() {

	userStorage := storage.New("data/users.json")
	userHandler := handlers.NewUserHandler(userStorage)

	mux := http.NewServeMux()

    mux.HandleFunc("GET /users",userHandler.GetUsers)
    mux.HandleFunc("POST /users",userHandler.CreateUser)
    mux.HandleFunc("GET /users/{id}",userHandler.GetUserByID)
    mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)

	logger.Init(true)
	http.ListenAndServe(":8080", mux)
}
