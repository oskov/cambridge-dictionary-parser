package parser

type Logger interface {
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
}

type NullLogger struct{}

func (l *NullLogger) Debug(msg string, args ...any) {}

func (l *NullLogger) Error(msg string, args ...any) {}
