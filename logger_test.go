package zlog

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	New(WithServiceName("test"),
		WithStdLevel(DebugLevel),
		WithFormatter(&JsonFormatter{}),
		WithOutputPath("./logs/", "app.log"),
		WithCleaner(&Clean{
			Interval: 24 * time.Hour,
			Reserve:  7 * 24 * time.Hour,
		}))
	Error("123123")
}
