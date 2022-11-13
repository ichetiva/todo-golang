package main

import (
	"time"

	"github.com/ichetiva/todo-golang/internal/use_cases/server"
)

func main() {
	server := server.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.Run()
}
