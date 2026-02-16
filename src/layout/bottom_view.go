package layout

import "github.com/rivo/tview"

type BottomView struct {
	Layout *tview.Flex
	TView  *tview.TextView
}

func NewBottomView() *BottomView {
	bv := &BottomView{
		TView: tview.NewTextView().SetTextAlign(tview.AlignLeft),
	}

	bv.TView.SetBorder(true)
	return bv
}

func (bv *BottomView) SetText(text string) *tview.TextView {
	bv.TView.SetText(text)
	return bv.TView
}
