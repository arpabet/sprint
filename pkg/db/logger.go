/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package db

import (
	"github.com/arpabet/template-server/pkg/constants"
	"github.com/dgraph-io/badger/v2"
	"log"
)

type loggerAdapter struct {
}

func (t* loggerAdapter) Errorf(format string, args ...interface{}) {
	log.Printf("ERROR " + format, args...)
}

func (t* loggerAdapter) Warningf(format string, args ...interface{}) {
	log.Printf("WARN " + format, args...)
}

func (t* loggerAdapter) Infof(format string, args ...interface{}) {
	log.Printf("INFO " + format, args...)
}

func (t* loggerAdapter) Debugf(format string, args ...interface{}) {
	if constants.IsDebug() {
		log.Printf(format, args...)
	}
}

func NewDBLogger() badger.Logger {
	return &loggerAdapter{}
}