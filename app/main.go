package main

import (
	"context"
	"log"
	"os"

	"github.com/g-chicken/mah-jong/app/controller/server"
	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/gateway/device"
	"github.com/g-chicken/mah-jong/app/gateway/rdb"
	"github.com/g-chicken/mah-jong/app/logger"
	"github.com/g-chicken/mah-jong/app/usecase"
	_ "github.com/go-sql-driver/mysql"
)

const (
	successCode = 0
	errorCode   = 1
)

func main() {
	grpcPort, closeFunc, err := initialize()
	if err != nil {
		closeFunc()
		log.Printf("fail to initialize (error = %v)", err)

		os.Exit(errorCode)
	}

	os.Exit(run(grpcPort, closeFunc))
}

func initialize() (int, func(), error) {
	// config
	config, err := getConfig()
	if err != nil {
		return config.GetGRPCPort(), func() {}, err
	}

	// logger
	if err := logger.SetLogger(); err != nil {
		return config.GetGRPCPort(), func() { logger.CloseLogger() }, err
	}

	// repository
	closeFunc, err := setRepositories(config)
	if err != nil {
		return config.GetGRPCPort(), func() {
			logger.CloseLogger()
			closeFunc()
		}, err
	}

	return config.GetGRPCPort(), func() {
		logger.CloseLogger()
		closeFunc()
	}, nil
}

func getConfig() (*domain.Config, error) {
	configRepo := device.NewConfigRepository()
	configUC := usecase.NewConfigUsecase(configRepo)

	return configUC.GetConfig(context.Background())
}

func setRepositories(config *domain.Config) (func(), error) {
	rdbGetterRepository, closeFunc, err := rdb.NewRDBGetterRepository(config)
	if err != nil {
		return closeFunc, err
	}

	playerRepo := rdb.NewPlayerRepository(rdbGetterRepository)

	domain.SetRepositories(playerRepo)

	return closeFunc, nil
}

func run(grpcPort int, closeFunc func()) int {
	defer closeFunc()

	srv := server.NewServer(grpcPort)

	if err := srv.Run(); err != nil {
		return errorCode
	}

	return successCode
}
