/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package run

import (
	"fmt"
	"github.com/arpabet/template-server/pkg/constants"
	"github.com/arpabet/template-server/pkg/db"
	"github.com/arpabet/template-server/pkg/util"
	"github.com/arpabet/template-server/pkg/client"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/gin-gonic/gin"
)

func ServerRun(daemon bool, masterKey string) error {
	fmt.Printf("Run Server %s Version %s\n", constants.ApplicationName, constants.GetAppInfo().Version)

	storage, err := db.NewStorage(constants.GetDatabaseFolder(), masterKey)
	if err != nil {
		return err
	}
	defer storage.Close()

	startTime := time.Now()

	var r  *gin.Engine
	var logWriter io.Writer
	if daemon {

		gin.SetMode(gin.ReleaseMode)
		r = gin.Default()
		// enable file logging if running as daemon

		logFile, err := os.OpenFile(constants.ExecutableLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		defer logFile.Close()

		logWriter = logFile
		r.Use(gin.LoggerWithWriter(logWriter))

		log.SetPrefix(constants.ApplicationName + ": ")
		log.SetFlags(0)
		log.SetOutput(logWriter)

	} else {
		logWriter = os.Stdout
		r = gin.Default()
	}

	signalChain := make(chan os.Signal, 1)
	signal.Notify(signalChain, os.Interrupt)

	api := r.Group("/api")

	api.GET("/status", func(c *gin.Context) {
		appInfo := constants.GetAppInfo()
		c.JSON(200, &client.StatusInfo{
			Version: appInfo.Version,
			Build:   appInfo.Build,
			Started: startTime.String(),
			MasterKeyHash:  util.GetKeyHash(masterKey),
		})
	})

	tlsConfig, err := util.LoadServerConfig(storage)
	if err != nil {
		log.Printf("ERROR: SSL certificates did not find in database.\n")
		return err
	}

	server := &http.Server{
		Addr:      constants.GetAddress(),
		TLSConfig: tlsConfig,
		Handler:   r}

	go func() {
		for _ = range signalChain {
			log.Printf("Stopping server\n")
			server.Close()
		}
	}()

	return server.ListenAndServeTLS("", "")

}
