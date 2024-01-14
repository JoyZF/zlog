package zlog

import (
	"fmt"
	"io"
	"os"
	"sync"
	"unsafe"
)

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

const (
	defaultServiceName = "zlog"
	defaultOutputPath  = "./logs/"
	defaultFileName    = "app.log"
)

var (
	log *logger
)

func New(opts ...Option) *logger {
	if log != nil {
		return log
	}
	if opts == nil {
		opts = append(opts, WithServiceName(defaultServiceName),
			WithOutputPath(defaultOutputPath, defaultFileName),
			WithFormatter(&TextFormatter{}),
			WithLevel(DebugLevel),
			WithCleaner(&Clean{}))
	}
	log = &logger{opt: initOptions(opts...)}
	log.entryPool = &sync.Pool{New: func() interface{} { return entry(log) }}
	if log.opt.cleaner != nil {
		log.opt.cleaner.Division()
	}
	return log
}

func StdLogger() *logger {
	return std
}

func SetOptions(opts ...Option) {
	std.SetOptions(opts...)
}

func (l *logger) SetOptions(opts ...Option) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, opt := range opts {
		opt(l.opt)
	}
}

func (l *logger) Writer() io.Writer {
	return l
}

func (l *logger) Write(data []byte) (int, error) {
	l.entry().write(l.opt.stdLevel, FmtEmptySeparate, *(*string)(unsafe.Pointer(&data)))
	return 0, nil
}

func (l *logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

func (l *logger) Debug(args ...interface{}) {
	l.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.entry().write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func (l *logger) Fatal(args ...interface{}) {
	l.entry().write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry().write(DebugLevel, format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.entry().write(InfoLevel, format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.entry().write(WarnLevel, format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.entry().write(ErrorLevel, format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.entry().write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.entry().write(FatalLevel, format, args...)
	os.Exit(1)
}
