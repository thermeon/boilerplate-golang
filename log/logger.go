package log

import (
	"context"
)

var inst Logger

type Logger interface {
	Debug(ctx context.Context, format string, args ...interface{})
	Warning(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, format string, args ...interface{})
	Fatal(ctx context.Context, format string, args ...interface{})
	Panic(ctx context.Context, format string, args ...interface{})
	Name() string
}

func Debug(ctx context.Context, format string, args ...interface{}) {
	inst.Debug(ctx, format, args...)
}

func Warning(ctx context.Context, format string, args ...interface{}) {
	inst.Warning(ctx, format, args...)
}

func Info(ctx context.Context, format string, args ...interface{}) {
	inst.Info(ctx, format, args...)
}

func Error(ctx context.Context, format string, args ...interface{}) {
	inst.Error(ctx, format, args...)
}

func Fatal(ctx context.Context, format string, args ...interface{}) {
	inst.Fatal(ctx, format, args...)
}

func Panic(ctx context.Context, format string, args ...interface{}) {
	inst.Panic(ctx, format, args...)
}
