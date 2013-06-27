package gotrail

import (
	logp "log"
	"io"
)

type Logger struct {
	*logp.Logger
	Debug bool
}

func New(out io.Writer) *Logger {
	return NewNamed(out, "")
}

func NewNamed(out io.Writer, prefix string) *Logger {
	logger := logp.New(out, "["+prefix+"] ", logp.LstdFlags)
	return &Logger{logger, false}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.Debug {
		l.Logger.Printf("DEBUG: "+format, args...)
	}
}

func (l *Logger) Debugln(args ...interface{}) {
	if l.Debug {
		args = append([]interface{}{"DEBUG:"}, args...)
		l.Logger.Println(args)
	}
}