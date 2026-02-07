package layout

import "github.com/rivo/tview"

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
		AddItem(tv.InputField, 1, 0, false)

	return tv
}
