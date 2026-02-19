package qrclogs

import (
	"log"

	"go.uber.org/zap"
)

func CreateLogger() *zap.Logger {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	return logger

}
