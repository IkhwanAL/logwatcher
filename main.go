package main

import (
	"context"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/ikhwanal/log_go/src"
	"github.com/ikhwanal/log_go/src/layout"
	"github.com/ikhwanal/log_go/src/logger"
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

	state := src.NewState()

	app := tview.NewApplication()

	topView := layout.NewTopView(&state.Search)
	logContent := layout.NewLogView()
	botView := layout.NewBottomView()

	flexGlobal := tview.NewFlex().SetDirection(tview.FlexRow)
	flexGlobal.AddItem(topView.Layout, 3, 0, false)
	flexGlobal.AddItem(logContent.Layout, 0, 4, false)
	flexGlobal.AddItem(botView.TView, 3, 0, false)

	botView.SetText("Command Mode")

	draw := src.NewRenderLog(app, logContent.TView, state)

	topView.SetChangeFunc(func() {
		draw.FilterContentAndDraw()
	})

	ctx, cancel := context.WithCancel(context.Background())

	// Capture Keystroke to Change Mode
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 's': // Search Mode
			if state.FocusOn == src.Search {
				break
			}
			state.FocusOn = src.Search
			app.SetFocus(topView.InputField)
			botView.SetText("Search Mode")
			return nil
		case 'v': // View Mode
			if state.FocusOn == src.ViewLog {
				break
			}
			state.FocusOn = src.ViewLog
			app.SetFocus(logContent.TView)
			botView.SetText("View Mode")
			return nil
		case 'c': // Command Mode
			if state.FocusOn == src.Command {
				break
			}
			state.FocusOn = src.Command
			app.SetFocus(nil)
			botView.SetText("Command Mode")
			return nil
		case 'q':
			cancel()
			app.Stop()
			return nil
		}

		if event.Key() == tcell.KeyEsc {
			if state.FocusOn == src.Command {
				return event
			}
			state.FocusOn = src.Command
			app.SetFocus(nil)
			botView.SetText("Command Mode")
			return nil
		}

		return event
	})

	watch := logger.WatchJournal(ctx)

	go logger.Pipe(watch, state, draw)

	if err := app.SetRoot(flexGlobal, true).SetFocus(flexGlobal).Run(); err != nil {
		panic(err)
	}
}
