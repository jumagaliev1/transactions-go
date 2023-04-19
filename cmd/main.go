package main

import (
	"context"
	"log"
	"transactions/internal/config"
	"transactions/internal/logger"
	"transactions/internal/service"
	"transactions/internal/storage"
	"transactions/internal/transport"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	cfg, err := config.New("configs/")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := storage.New(ctx, cfg, logger.RequestLogger{})
	if err != nil {
		log.Fatal(err)
	}

	serv, err := service.New(repo, *cfg, logger.RequestLogger{})
	if err != nil {
		log.Fatal(err)
	}

	if err := transport.Run(serv); err != nil {
		log.Fatal(err)
	}
}
