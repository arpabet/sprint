/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package service

import (
	"github.com/arpabet/timeuuid"
	"github.com/arpabet/templateserv/pkg/app"
	"github.com/arpabet/templateserv/pkg/util"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"sync"
	"time"
	"errors"
)

type nodeService struct {
	ConfigService    app.ConfigService  `inject`
	Log              *zap.Logger        `inject`

	initOnce         sync.Once

	nodeIdHex        string
	nodeId           uint64

	lastTimestamp    atomic.Int64
	clock            atomic.Int32
}

func NodeService() app.NodeService {
	return &nodeService{}
}

func (t *nodeService) NodeId() uint64 {
	t.init()
	return t.nodeId
}


func (t *nodeService) NodeIdHex() string {
	t.init()
	return t.nodeIdHex
}

func (t *nodeService) init() {
	t.initOnce.Do(func() {
		if err := t.doInit(); err != nil {
			t.Log.Error("NodeId Init", zap.Error(err))
		}
	})
}

func (t *nodeService) doInit() error {

	var err error
	t.nodeIdHex, err = t.ConfigService.Get(app.NodeId)
	if err != nil {
		return err
	}
	if t.nodeIdHex == "" {
		return errors.New("Empty NodeId in Config")
	}
	t.nodeId, err = util.ParseNodeId(t.nodeIdHex)
	if err != nil {
		return err
	}

	return nil
}

func (t *nodeService) Issue() timeuuid.UUID {

	uuid := timeuuid.NewUUID(timeuuid.TimebasedVer1)
	uuid.SetTime(time.Now())
	uuid.SetNode(int64(t.nodeId))

	for {

		curr := uuid.UnixTime100Nanos()
		old := t.lastTimestamp.Load()
		if old == curr {
			uuid.SetClockSequence(int(t.clock.Inc()))
			break
		}

		if t.lastTimestamp.CAS(old, curr) {
			t.clock.Store(0)
			break
		}

		old = t.lastTimestamp.Load()
		if old > curr {
			uuid.SetTime(time.Now())
		}

	}

	return uuid

}

func (t *nodeService) Parse(uuid timeuuid.UUID) (timestampMillis int64, nodeId int64, clock int) {
	return uuid.UnixTimeMillis(), uuid.Node(), uuid.ClockSequence()
}

