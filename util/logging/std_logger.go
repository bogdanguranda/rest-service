package logging

import (
	"fmt"
	"log"
	"os"
)

type StdLogger struct {
	globalLevel LogLevel
}

func NewStdLogger(globalLevel LogLevel) *StdLogger {
	log.SetOutput(os.Stdout)
	return &StdLogger{globalLevel: globalLevel}
}

func (sl StdLogger) Log(level LogLevel, message string) {
	if !sl.isLevelEnough(level) {
		return
	}
	log.SetPrefix(fmt.Sprintf("%s: ", level))
	log.Println(message)
}

func (sl StdLogger) isLevelEnough(level LogLevel) bool {
	if sl.globalLevel == LogLevelDebug {
		return true
	}

	if sl.globalLevel == LogLevelInfo && level != LogLevelDebug {
		return true
	}

	if sl.globalLevel == LogLevelError && level == LogLevelError {
		return true
	}

	return false
}
