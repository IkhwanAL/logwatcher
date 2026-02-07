package log

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/rivo/tview"
)

func Pipe(out io.ReadCloser, app *tview.Application, tv *tview.TextView) {
	scanner := bufio.NewScanner(out)

	for scanner.Scan() {
		line := scanner.Text()

		app.QueueUpdateDraw(func() {
			fmt.Fprint(tv, line+"\n")
		})
	}

	if err := scanner.Err(); err != nil {
		log.Println("JournalCTL failed to scan: ", err.Error())
	}
}
