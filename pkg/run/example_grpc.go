/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package run

import (
	c "context"
	"github.com/arpabet/sprint/pkg/pb"
	"time"
)

func (t *serverImpl) Status(ctx c.Context, request *pb.StatusRequest) (*pb.StatusResponse, error) {

	uptime := time.Now().Sub(t.startTime) / time.Millisecond

	resp := &pb.StatusResponse{
		NodeId:   int64(t.NodeService.NodeId()),
		Uptime:   int64(uptime),
	}

	return resp, nil
}

