package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BuildLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zapcore.InfoLevel)

	var err error
	logger, err := config.Build()
	if err != nil {
		panic("Failed to setup logger")
	}

	return logger
}
