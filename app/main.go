package main

import (
	"context"
	"log"
	"os"

	"github.com/g-chicken/mah-jong/app/controller/server"
	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/device"
	"github.com/g-chicken/mah-jong/app/logger"
	"github.com/g-chicken/mah-jong/app/usecase"
)

const (
	successCode = 0
	errorCode   = 1
)

func main() {
	configRepo := device.NewConfigRepository()
	configUC := usecase.NewConfigUsecase(configRepo)

	config, err := configUC.GetConfig(context.Background())
	if err != nil {
		log.Printf("fail to build config (error = %v)\n", err)

		os.Exit(1)
	}

	os.Exit(run(config))
}

func run(config *domain.Config) int {
	err := logger.SetLogger()
	defer logger.CloseLogger()

	if err != nil {
		log.Printf("fail to build logger (error = %v)\n", err)

		return errorCode
	}

	// domain.SetRepositories(nil)

	srv := server.NewServer(config)

	if err := srv.Run(); err != nil {
		return errorCode
	}

	return successCode
}
