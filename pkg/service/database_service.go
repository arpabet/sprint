/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package service

import (
	"fmt"
	"github.com/arpabet/sprint/pkg/app"
	"github.com/arpabet/sprint/pkg/pb"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
)

/**
	Database Service Impl

	Alex Shvid
*/


type databaseService struct {
	ConfigService    app.ConfigService  `inject`
	Storage          app.Storage        `inject`
	Log              *zap.Logger        `inject`
}

func DatabaseService() app.DatabaseService {
	return &databaseService{}
}

func (t *databaseService) Execute(query string, cb func(app.Record) bool) error {

	query = strings.Trim(query, "")
	if strings.HasSuffix(query, ";") {
		query = query[:len(query)-1]
	}

	cmdEnd := strings.Index(query, " ")
	if cmdEnd == -1 {
		cmdEnd = len(query)
	}

	cmd := query[:cmdEnd]
	args := strings.Trim(query[cmdEnd:], " ")

	switch cmd {
	case "get":
		key := []byte(args)
		if value, err := t.Storage.Get([]byte(args), false); err != nil {
			return err
		} else {
			cb(app.Record{
				Key:   key,
				Value: value,
			})
		}
	default:
		return errors.Errorf("unknown command cmd=%s, args=%s", cmd, args)
	}

	return nil
}

func (t *databaseService) Console(stream app.DatabaseConsoleStream) error {

	for {
		request, err := stream.Recv()
		if err != nil {
			break
		}

		err = t.Execute(request.Query, func(record app.Record) bool {

			rec := &pb.DatabaseResponse{
				Status:  200,
				Content: []byte("content"),
			}

			return stream.Send(rec) == nil

		})

		if err != nil {
			err = stream.Send(&pb.DatabaseResponse{
				Status:    501,
				Content:   []byte(fmt.Sprintf("exec error, %v", err)),
			})
		}

		err = stream.Send(&pb.DatabaseResponse{
			Status:  100,
		})

		if err != nil {
			return err
		}

	}

	return nil
}


