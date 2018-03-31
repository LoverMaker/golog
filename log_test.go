package golog

import (
	"testing"
	"github.com/LoverMaker/golog/output"
	"github.com/LoverMaker/golog/formatter"
)

func TestLog(t *testing.T) {
	Register(output.NewFileLogger("test", "info", &formatter.TextFormat{}))
	Info("test")
}
