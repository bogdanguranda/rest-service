package logging

type Logger interface {
	Log(level LogLevel, message string)
}
