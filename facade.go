package zlog

import (
	"context"
	"fmt"
	"os"
)

type trace struct {
	key   string
	value interface{}
}

var std = New()

// std logger
func Debug(args ...interface{}) {
	std.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func DebugWithCtx(ctx context.Context, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Debug(args)
}

func Info(args ...interface{}) {
	std.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func InfoWithCtx(ctx context.Context, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Info(args)
}

func Warn(args ...interface{}) {
	std.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func WarnWithCtx(ctx context.Context, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Warn(args)
}

func Error(args ...interface{}) {
	std.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func ErrorWithCtx(ctx context.Context, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Error(args)
}

func Panic(args ...interface{}) {
	std.entry().write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func PanicWithCtx(ctx context.Context, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Panic(args)
}

func Fatal(args ...interface{}) {
	std.entry().write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func FatalWithCtx(ctx context.Context, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Fatal(args)
}

func Debugf(format string, args ...interface{}) {
	std.entry().write(DebugLevel, format, args...)
}

func DebugfWithCtx(ctx context.Context, format string, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Debugf(format, args)
}

func Infof(format string, args ...interface{}) {
	std.entry().write(InfoLevel, format, args...)
}

func InfofWithCtx(ctx context.Context, format string, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Infof(format, args)
}

func Warnf(format string, args ...interface{}) {
	std.entry().write(WarnLevel, format, args...)
}

func WarnfWithCtx(ctx context.Context, format string, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Warnf(format, args)
}

func Errorf(format string, args ...interface{}) {
	std.entry().write(ErrorLevel, format, args...)
}

func ErrorfWithCtx(ctx context.Context, format string, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Errorf(format, args)
}

func Panicf(format string, args ...interface{}) {
	std.entry().write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func PanicfWithCtx(ctx context.Context, format string, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Panicf(format, args)
}

func Fatalf(format string, args ...interface{}) {
	std.entry().write(FatalLevel, format, args...)
	os.Exit(1)
}

func FatalfWithCtx(ctx context.Context, format string, args ...interface{}) {
	if len(std.opt.traceKey) != 0 {
		value := ctx.Value(std.opt.traceKey)
		args = append(args, trace{key: std.opt.traceKey, value: value})
	}
	Fatalf(format, args)
}
