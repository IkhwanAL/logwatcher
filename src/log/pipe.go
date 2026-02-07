package logger

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/rivo/tview"
)

// Pipe 'out' to the 'app' and show it to the 'view'
func Pipe(out io.Reader, app *tview.Application, view *tview.TextView) {
	scanner := bufio.NewScanner(out)

	for scanner.Scan() {
		line := scanner.Text()

		app.QueueUpdateDraw(func() {
			fmt.Fprint(view, line+"\n")
		})
	}

	if err := scanner.Err(); err != nil {
		log.Println("JournalCTL failed to scan: ", err.Error())
	}
}
