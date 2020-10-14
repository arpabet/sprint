/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package service

import (
	"fmt"
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/context"
	"github.com/arpabet/sprint/pkg/db"
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

	scan := asList(
		logger,
		storage,
		ConfigService(),
		NodeService(),
		DatabaseService())

	if app.Scan != nil {
		scan = append(scan, app.Scan)
	}

	ctx, err := context.Create(scan...)
	if err != nil {
		return nil, err
	}

	if app.Initialized != nil {
		return ctx, app.Initialized(ctx)
	} else {
		return ctx, nil
	}

}

func asList(list... interface{}) []interface{} {
	return list
}
