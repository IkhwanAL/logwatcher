package logger

import (
	"fmt"

	"github.com/rivo/tview"
)

type Log interface {
	Parse() string
}

func Pipe(ch <-chan Log, app *tview.Application, view *tview.TextView) {
	for watch := range ch {
		app.QueueUpdateDraw(func() {
			fmt.Fprint(view, watch.Parse()+"\n")
		})
	}
}
