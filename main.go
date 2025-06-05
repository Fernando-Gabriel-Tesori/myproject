package main

import (
	"log"
	"myproject/api"
	"myproject/config"
	"myproject/pkg/database"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectMongo(cfg.DatabaseURL)

	log.Printf("🚀 Servidor rodando na porta %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, api.RegisterRoutes()))
}
