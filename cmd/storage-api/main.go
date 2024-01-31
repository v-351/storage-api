package main

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
	"github.com/v-351/storage-api/internal/app"
	"github.com/v-351/storage-api/internal/config"
)

// @title	Storage API
// @version 1.0
// @host	localhost:8080
func main() {
	var (
		config config.Config
		err    error
	)

	// Для запуска без экспорта .env в compose
	// err = godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = envconfig.Process(context.Background(), &config)
	if err != nil {
		log.Fatal(err)
	}

	app.Run(config)
}
