/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package run

import (
	c "context"
	"github.com/pkg/errors"
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/pb"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

/**
	ConfigService gRPC Impl
*/


func (t *serverImpl) Node(ctx c.Context, request *pb.NodeRequest) (*pb.NodeResponse, error) {

	uptime := time.Now().Sub(t.startTime).Milliseconds()

	resp := &pb.NodeResponse{
		ApplicationName: app.ApplicationName,
		Version:  app.Version,
		Build:    app.Build,
		NodeId:   int64(t.NodeService.NodeId()),
		Uptime:   uptime,
	}

	return resp, nil
}

func (t *serverImpl) Shutdown(ctx c.Context, req *pb.ShutdownRequest) (*pb.ShutdownResponse, error) {
	t.Log.Info("Received shutdown signal")
	t.restarting.Store(req.Restart)
	time.AfterFunc(app.StopDelay, func() {
		t.signalChain <- os.Interrupt
	})
	return new(pb.ShutdownResponse), nil
}

func (t *serverImpl) SetConfig(ctx c.Context, request *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	if request.Key == "" {
		return nil, errors.Errorf("empty config key")
	}
	if err := t.ConfigService.Set(request.Key, request.Value); err != nil {
		return nil, err
	}
	return new(pb.SetConfigResponse), nil
}

func (t *serverImpl) GetConfig(ctx c.Context, request *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	if request.Key == "" {
		return nil, errors.Errorf("empty config key")
	}
	entry, err := t.getConfigImpl(request.Key)
	if err != nil {
		return nil, err
	}
	resp := &pb.GetConfigResponse{ Entry : entry }
	return resp, nil
}

func (t *serverImpl) Configuration(ctx c.Context, request *pb.ConfigurationRequest) (*pb.ConfigurationResponse, error) {

	var entries []*pb.ConfigEntry
	t.ConfigService.GetAll(func(key, value string) {
		if strings.Contains(key, "password") {
			value = "******"
		}
		entries = append(entries, &pb.ConfigEntry{
			Key:                  key,
			Value:                value,
		})
	})
	resp := &pb.ConfigurationResponse{ Entry : entries }
	return resp, nil

}

func (t *serverImpl) getConfigImpl(key string) (*pb.ConfigEntry, error) {
	var err error
	var value string
	if value, err = t.ConfigService.Get(key); err != nil {
		return nil, err
	}
	if strings.Contains(key, "password") {
		value = "******"
	}
	return &pb.ConfigEntry{
		Key:                  key,
		Value:                value,
	}, nil

}

func (t *serverImpl) DatabaseConsole(stream pb.NodeService_DatabaseConsoleServer) error {

	defer func() {
		if r := recover(); r != nil {
			t.Log.Error("DatabaseConsoleRecover",
				zap.Any("recover", r))
		}
	}()

	err := t.DatabaseService.Console(stream)
	if err != nil {
		t.Log.Error("DatabaseConsole",
			zap.Error(err))
	}
	return err
}



