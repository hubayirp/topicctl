package zk

import (
	szk "github.com/samuel/go-zookeeper/zk"
	log "github.com/sirupsen/logrus"
)

// ZKDebugLogger is a logger that satisfies the szk.Logger interface.
type ZKDebugLogger struct{}

var _ szk.Logger = (*ZKDebugLogger)(nil)

// Printf sends samuel zk log messages to logrus at the debug level.
func (l *ZKDebugLogger) Printf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
