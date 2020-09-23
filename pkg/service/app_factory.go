/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package service

import (
	"fmt"
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/arpabet/context"
	"github.com/arpabet/templateserv/pkg/db"
	"go.uber.org/zap"
)

func newLogger() (*zap.Logger, error) {

	logFile := app.GetLogFile()
	fmt.Printf("Write to log: %s\n", logFile)

	if app.IsDaemon() {

		cfg := zap.NewDevelopmentConfig()
		cfg.OutputPaths = []string{
			logFile,
		}
		return cfg.Build()

	} else {
		return zap.NewDevelopment()
	}

}

func CreateContext(masterKey string) (context.Context, error) {

	logger, err := newLogger()
	if err != nil {
		return nil, err
	}

	storage, err := db.NewStorage(app.GetDataFolder(), masterKey)
	if err != nil {
		return nil, err
	}

	return context.Create(
		logger,
		storage,
		ConfigService(),
		NodeService())

}
