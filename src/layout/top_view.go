package layout

import (
	"sync"
	"time"

	"github.com/rivo/tview"
)

type SearchState string

type TopView struct {
	Layout     *tview.Flex
	TView      *tview.TextView
	InputField *tview.InputField
	Search     *string
}

func NewTopView(searhRef *string) *TopView {
	tv := &TopView{
		Layout: tview.NewFlex().SetDirection(tview.FlexRow),
		TView: tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText("COMMAND"),
		InputField: tview.NewInputField().SetLabel("Search: "),
		Search:     searhRef,
	}

	tv.Layout.
		AddItem(tv.TView, 2, 0, false).
		AddItem(tv.InputField, 1, 0, false)

	return tv
}

func (tv *TopView) SetChangeFunc(event func()) {
	debounceSearch := NewDebounce(func() {
		val := tv.InputField.GetText()
		*tv.Search = val
		event()
	}, 300*time.Millisecond)

	tv.InputField.SetChangedFunc(func(text string) {
		debounceSearch()
	})
}

func NewDebounce(event func(), delay time.Duration) func() {
	var timer *time.Timer // a State to Remember
	var mu sync.Mutex     // To Make Sure it run one at the time if there are multiple debounce

	return func() {
		mu.Lock()
		defer mu.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(delay, event)
	}
}
