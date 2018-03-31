package output

import (
"io"
"github.com/sirupsen/logrus"
)

type Interface interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	SetLevel(level logrus.Level)
	Reload(w io.WriteCloser) error
	Close() error
}
