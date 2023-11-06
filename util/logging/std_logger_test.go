package logging

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestLogger struct {
	output *bytes.Buffer
}

func NewTestLogger() *TestLogger {
	return &TestLogger{output: &bytes.Buffer{}}
}

func (tl *TestLogger) Log(level LogLevel, message string) {
	if !tl.isLevelEnough(level) {
		return
	}
	tl.output.WriteString(fmt.Sprintf("%s: %s\n", level, message))
}

func (tl *TestLogger) isLevelEnough(level LogLevel) bool {
	return true // TestLogger captures all log levels for testing
}

func (tl *TestLogger) GetOutput() string {
	return tl.output.String()
}

func TestDebugLogging(t *testing.T) {
	testLogger := NewTestLogger()
	logger := NewStdLogger(LogLevelDebug)

	log.SetOutput(testLogger.output)
	logger.Log(LogLevelDebug, "Debug message") // This should be logged
	logger.Log(LogLevelInfo, "Info message")   // This should be logged
	logger.Log(LogLevelError, "Error message") // This should be logged

	actualDebugOutput := testLogger.GetOutput()
	assert.Contains(t, actualDebugOutput, "Debug message", "Debug log output mismatch")
	assert.Contains(t, actualDebugOutput, "Info message", "Debug log output mismatch")
	assert.Contains(t, actualDebugOutput, "Error message", "Debug log output mismatch")
}

func TestInfoLogging(t *testing.T) {
	testLogger := NewTestLogger()
	logger := NewStdLogger(LogLevelInfo)

	log.SetOutput(testLogger.output)
	logger.Log(LogLevelDebug, "Debug message") // This should not be logged
	logger.Log(LogLevelInfo, "Info message")   // This should be logged
	logger.Log(LogLevelError, "Error message") // This should be logged

	actualDebugOutput := testLogger.GetOutput()
	assert.NotContains(t, actualDebugOutput, "Debug message", "Debug log output mismatch")
	assert.Contains(t, actualDebugOutput, "Info message", "Debug log output mismatch")
	assert.Contains(t, actualDebugOutput, "Error message", "Debug log output mismatch")
}

func TestErrorLogging(t *testing.T) {
	testLogger := NewTestLogger()
	logger := NewStdLogger(LogLevelError)

	log.SetOutput(testLogger.output)
	logger.Log(LogLevelDebug, "Debug message") // This should not be logged
	logger.Log(LogLevelInfo, "Info message")   // This should not be logged
	logger.Log(LogLevelError, "Error message") // This should be logged

	actualDebugOutput := testLogger.GetOutput()
	assert.NotContains(t, actualDebugOutput, "Debug message", "Debug log output mismatch")
	assert.NotContains(t, actualDebugOutput, "Info message", "Debug log output mismatch")
	assert.Contains(t, actualDebugOutput, "Error message", "Debug log output mismatch")
}
