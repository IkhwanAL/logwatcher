package test

import (
	"testing"
	"time"

	"github.com/ikhwanal/log_go/src/layout"
)

func TestDebounce(t *testing.T) {
	var done = make(chan struct{})

	aDebounceCall := func() {
		close(done)
	}

	debounce := layout.NewDebounce(aDebounceCall, 300*time.Millisecond)

	debounce()
	debounce()
	debounce()

	select {
	case <-done:
		return
	case <-time.After(1 * time.Second):
		t.Errorf("debounce never fired")
	}
}
