package main

import (
	"log"
	"net/http"

	"myproject/api"
	"myproject/config"
	"myproject/internal/repository"
	"myproject/pkg/database"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectMongo(cfg.DatabaseURL)
	repository.InitUserRepository()

	log.Printf("🚀 Servidor rodando na porta %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, api.RegisterRoutes()))
}
