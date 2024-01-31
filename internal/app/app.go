package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/v-351/storage-api/internal/config"
	"github.com/v-351/storage-api/internal/controller/http"
	"github.com/v-351/storage-api/internal/repository/postgres"
	"github.com/v-351/storage-api/internal/usecase"
)

func Run(config config.Config) {
	storage := postgres.New(config)
	defer storage.Close()

	usecase := usecase.New(storage)

	server := http.New(config, usecase)
	go server.Start(config.Http.Address)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)
	<-shutdown
	server.Shoutdown()
}
