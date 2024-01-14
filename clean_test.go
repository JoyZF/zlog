package zlog

import (
	"testing"
	"time"
)

func TestClean_do(t *testing.T) {
	clean := Clean{Interval: 10 * time.Second, Reserve: 1 * time.Second}
	clean.clean()
	for {

	}
}
