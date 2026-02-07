package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/ikhwanal/log_go/src/layout"
	logger "github.com/ikhwanal/log_go/src/log"
	"github.com/rivo/tview"
)

func main() {
	// Log To file
	// since the screen will be occupied by UI
	outLog, err := os.OpenFile("stdout.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Log Failed To Open")
	}
	defer outLog.Close()

	log.SetOutput(outLog)

	state := NewState()

	app := tview.NewApplication()

	flexGlobal := tview.NewFlex().SetDirection(tview.FlexRow)

	topView := layout.NewTopView()
	flexGlobal.AddItem(topView.Layout, 3, 0, false)

	logContent := layout.NewLogView()
	flexGlobal.AddItem(logContent.Layout, 0, 4, false)

	// Capture Keystroke to Change Mode
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 's': // Search Mode
			if state.FocusOn == Search {
				break
			}
			state.FocusOn = Search
			app.SetFocus(topView.InputField)
			return nil
		case 'v':
			if state.FocusOn == ViewLog {
				break
			}
			state.FocusOn = ViewLog
			app.SetFocus(logContent.TView)
			return nil
		case 'c':
			if state.FocusOn == Command {
				break
			}
			state.FocusOn = Command
			app.SetFocus(nil)
			return nil
		}

		if event.Key() == tcell.KeyEsc {
			app.Stop()
			return nil
		}

		return event
	})

	out := logger.WatchJournal()
	go logger.Pipe(out, app, logContent.TView)

	if err := app.SetRoot(flexGlobal, true).SetFocus(flexGlobal).Run(); err != nil {
		panic(err)
	}
}
