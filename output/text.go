package output

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"github.com/LoverMaker/golog/formatter"
)
type TextLogger struct {
	w io.WriteCloser
}

func newTextLogger(w io.WriteCloser, level logrus.Level, format logrus.Formatter) *TextLogger {
	logger := &TextLogger{
		w: w,
	}
	logrus.SetOutput(w)
	logrus.SetLevel(level)
	logrus.SetFormatter(format)

	return logger
}

func NewFileLogger(logPath string, logLevel string, format logrus.Formatter) *TextLogger    {
	var w io.WriteCloser
	if logPath != "" {
		var err error
		w, err = os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
		if err != nil {
			panic("error to create log file")
		}
	} else {
		w = os.Stdout
	}

	return newTextLogger(
		w,
		ParseLevel(logLevel),
		&formatter.TextFormat{},
	)
}

func NewConsoleLogger(logLevel string, format logrus.Formatter) *TextLogger  {
	return newTextLogger(
		os.Stdout,
		ParseLevel(logLevel),
		&formatter.TextFormat{},
	)
}

func (t *TextLogger) SetLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

func (t *TextLogger) Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func (t *TextLogger) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (t *TextLogger) Info(args ...interface{}) {
	logrus.Info(args...)
}

func (t *TextLogger) Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (t *TextLogger) Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func (t *TextLogger) Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func (t *TextLogger) Error(args ...interface{}) {
	logrus.Error(args...)
}

func (t *TextLogger) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func (t *TextLogger) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func (t *TextLogger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func (t *TextLogger) Close() error {
	return t.w.Close()
}

func (t *TextLogger) Reload(w io.WriteCloser) error {
	tw := t.w
	logrus.SetOutput(w)

	return tw.Close()
}
