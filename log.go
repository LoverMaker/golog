package golog

import (
	"io"
	"fmt"
	"github.com/LoverMaker/golog/output"
	"github.com/LoverMaker/golog/formatter"
)

type Logger struct {
	output output.Interface
}

var DefaultLogger Logger

func init() {
	fmt.Println("init log")
	DefaultLogger.output = output.NewConsoleLogger("info", &formatter.TextFormat{})
}

func Register(output output.Interface) {
	DefaultLogger.output = output
}

func Debug(args ...interface{}) {
	DefaultLogger.output.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	DefaultLogger.output.Debugf(format, args...)
}

func Info(args ...interface{}) {
	DefaultLogger.output.Info(args...)
}

func Infof(format string, args ...interface{}) {
	DefaultLogger.output.Infof(format, args...)
}

func Warn(args ...interface{}) {
	DefaultLogger.output.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	DefaultLogger.output.Warnf(format, args...)
}

func Error(args ...interface{}) {
	DefaultLogger.output.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	DefaultLogger.output.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	DefaultLogger.output.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	DefaultLogger.output.Fatalf(format, args...)
}

func SetLevel(level string) {
	DefaultLogger.output.SetLevel(output.ParseLevel(level))
}

func Reload(w io.WriteCloser) error {
	return DefaultLogger.output.Reload(w)
}

func Close() error {
	return DefaultLogger.output.Close()
}
