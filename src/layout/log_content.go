package layout

import "github.com/rivo/tview"

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
	lv.Layout.AddItem(lv.TView, 0, 1, true)

	return lv
}
