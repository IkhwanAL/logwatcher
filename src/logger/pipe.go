package logger

import (
	"github.com/ikhwanal/log_go/src"
)

type Log interface {
	Parse() string
}

// This is Where Print Everthing Without Expection and Watch Any Channel Log
func Pipe(ch <-chan Log, state *src.State, draw *src.RenderLog) {
	for watch := range ch {
		text := watch.Parse()
		state.Sync.Lock()
		state.Contents = append(state.Contents, text)
		state.Sync.Unlock()

		if state.Search == "" {
			draw.SetContentAndDraw()
		}
	}
}
