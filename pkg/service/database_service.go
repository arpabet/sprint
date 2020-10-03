/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package service

import (
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/arpabet/templateserv/pkg/pb"
	"go.uber.org/zap"
	"fmt"
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


