/*
* Copyright 2020-present Arpabet Inc. All rights reserved.
 */

package run

import (
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
	"time"
)

func (t *serverImpl) Autoupdate(distrFile string) error {

	var err error
	t.autoupdate, err = fsnotify.NewWatcher()
	if err != nil {
		return errors.Errorf("fail to create watcher %v", err)
	}

	t.distrStat, err = os.Stat(distrFile)
	if err != nil {
		return errors.Errorf("distr file not found %v", err)
	}

	err = t.autoupdate.Add(distrFile)
	if err != nil {
		return err
	}

	t.Log.Info("Autoupdate AddWatch", zap.String("distrFile", distrFile))
	go t.AutoupdateLoop()

	return nil

}

func (t *serverImpl) AutoupdateLoop() {
	for {
		select {
		case event, ok := <-t.autoupdate.Events:
			if ok {
				t.AutoupdateEvent(event)
			}
		case err, ok := <-t.autoupdate.Errors:
			if ok {
				t.Log.Error("Autoupdate Watcher", zap.Error(err))
			}
		case <-t.autoupdateDone:
			t.Log.Info("Autoupdate Watcher shutdown")
			return
		}
	}
}

func (t *serverImpl) AutoupdateEvent(event fsnotify.Event) {
	if event.Op == fsnotify.Create || event.Op == fsnotify.Write {
		if stat, err := os.Stat(event.Name); err == nil {
			if t.distrStat.Name() != event.Name {
				t.Log.Error("Autoupdate Wrong Name",
					zap.String("eventName", event.Name),
					zap.String("distrName", t.distrStat.Name()),
				)
				return
			}
			if stat.Size() != t.distrStat.Size() || stat.ModTime().After(t.distrStat.ModTime()) {
				t.Log.Info("Autoupdate Detection",
					zap.String("eventName", event.Name),
					zap.String("op", event.Op.String()),
					zap.Int64("size", stat.Size()),
					zap.Time("mt", stat.ModTime()),
					zap.Int64("distr.size", t.distrStat.Size()),
					zap.Time("distr.mt", t.distrStat.ModTime()),
				)
				t.requestUpdate()
			} else {
				t.Log.Warn("Autoupdate NoChange",
					zap.String("eventName", event.Name),
					zap.String("op", event.Op.String()),
					zap.Int64("size", stat.Size()),
					zap.Time("mt", stat.ModTime()),
					zap.Int64("distr.size", t.distrStat.Size()),
					zap.Time("distr.mt", t.distrStat.ModTime()),
				)
			}
		}
	}
}

func (t *serverImpl) requestUpdate() {
	request := time.Now().UnixNano()
	t.requestUpdateTimestamp.Store(request)
	time.AfterFunc(time.Second, func() {
		if t.requestUpdateTimestamp.Load() == request {
			t.Log.Info("Autoupdate Restarting")
			t.restarting.Store(true)
			t.signalChain <- os.Interrupt
		}
	})
}

