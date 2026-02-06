package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/rivo/tview"
)

type TopView struct {
	Layout     *tview.Flex
	TView      *tview.TextView
	InputField *tview.InputField
}

func NewTopView() *TopView {
	tv := &TopView{
		Layout: tview.NewFlex().SetDirection(tview.FlexRow),
		TView: tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText("COMMAND"),
		InputField: tview.NewInputField().SetLabel("Search: "),
	}

	tv.Layout.
		AddItem(tv.TView, 2, 0, false).
		AddItem(tv.InputField, 1, 0, true)

	return tv
}

type LogContent struct {
	Layout *tview.Flex
	TView  *tview.TextView
}

func NewLogView() *LogContent {
	lv := &LogContent{
		Layout: tview.NewFlex().SetDirection(tview.FlexRow),
		TView: tview.NewTextView().
			SetDynamicColors(true).
			SetRegions(true).
			SetWrap(true).
			SetTextAlign(tview.AlignLeft),
	}

	lv.TView.SetBorder(true).SetTitle("LOG")
	lv.Layout.AddItem(lv.TView, 0, 1, false)

	return lv
}

func PipeLogJournalCTL(app *tview.Application, tv *tview.TextView) {
	cmd := exec.Command("journalctl", "-f", "--output=short")
	out, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatalf("Failed To Pipe In To Log TUI ")
	}
	if err = cmd.Start(); err != nil {
		log.Fatalf("Failed To Start The Plumbing")
	}

	go func() {
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
	}()
}

func main() {
	outLog, err := os.OpenFile("stdout.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Log Failed To Open")
	}
	defer outLog.Close()

	log.SetOutput(outLog)

	app := tview.NewApplication()

	flexGlobal := tview.NewFlex().SetDirection(tview.FlexRow)

	topView := NewTopView()

	flexGlobal.AddItem(topView.Layout, 3, 0, true)

	logContent := NewLogView()

	flexGlobal.AddItem(logContent.Layout, 0, 4, false)

	PipeLogJournalCTL(app, logContent.TView)

	if err := app.SetRoot(flexGlobal, true).SetFocus(flexGlobal).Run(); err != nil {
		panic(err)
	}
}
