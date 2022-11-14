package main

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/ichetiva/todo-golang/config"
	"github.com/ichetiva/todo-golang/internal/use_cases/server"
)

func main() {
	server := server.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	configPath, err := GetConfigPath()
	if err != nil {
		log.Fatal(err)
	}
	config, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	server.Run(config)
}

func GetConfigPath() (string, error) {
	basePath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	configPath := path.Join(basePath, "config", "config.yml")
	return configPath, nil
}
