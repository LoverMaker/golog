package formatter

import (
	"bytes"
	"fmt"
	"runtime"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/inconshreveable/log15/stack"
)

const defaultTimestampFormat = "2006-01-02 15:04:05"

type TextFormat struct {
	// Disable timestamp logging. useful when output is redirected to logging
	// system that already adds timestamps.
	DisableTimestamp bool

	// Enable logging the full timestamp when a TTY is attached instead of just
	// the time passed since beginning of execution.
	FullTimestamp bool

	// TimestampFormat to use for display when a full timestamp is printed
	TimestampFormat string
}

func (t *TextFormat) Format(entry *logrus.Entry) ([]byte, error) {
	keys := make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}
	p := make([]uintptr, 7)
	count := runtime.Callers(3, p)
	call := stack.Call(p[count-2])
	//file = path.Base(file)

	b := &bytes.Buffer{}

	timestampFormat := t.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	b.WriteByte('[')

	if !t.DisableTimestamp {
		b.WriteString(time.Now().Format(timestampFormat))
		b.WriteString("] ")
	}
	b.WriteString(" ")
	b.WriteString(entry.Level.String())
	b.WriteString(fmt.Sprintf(" [%+v] > ", call))

	if entry.Message != "" {
		b.WriteString(entry.Message)
		b.WriteByte(' ')
	}

	for _, key := range keys {
		b.WriteString(key)
		b.WriteByte(' ')
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}
