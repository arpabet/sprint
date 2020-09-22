package run

import (
	c "context"
	"github.com/arpabet/templateserv/pkg/pb"
	"time"
)

func (t *serverImpl) Node(ctx c.Context, request *pb.NodeRequest) (*pb.NodeResponse, error) {

	uptime := time.Now().Sub(t.startTime) / time.Millisecond

	resp := &pb.NodeResponse{
		NodeId:   int64(t.NodeService.NodeId()),
		Uptime:   int64(uptime),
	}

	return resp, nil
}

