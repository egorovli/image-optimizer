package main

import (
	"go.uber.org/zap"
)

func mustInitLogger() {
	var unsugaredLogger *zap.Logger
	var err error

	if config.Env == "development" {
		unsugaredLogger, err = zap.NewDevelopment()
	} else {
		unsugaredLogger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = unsugaredLogger.Sync()
	}()

	logger = unsugaredLogger.Sugar()
}
