package src

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

// Here Lies a Command Allow Them To Render a Log
// So if there's involve sync or channel use this Struct to log
type RenderLog struct {
	App   *tview.Application // The Application
	View  *tview.TextView    // Log View
	State *State
}

func (r *RenderLog) Exec() {
	r.App.QueueUpdateDraw(func() {
		r.View.Clear()
		for _, content := range r.State.FilterContent {
			fmt.Fprint(r.View, content+"\n")
		}
	})
}

func (r *RenderLog) SetContentAndDraw() {
	r.State.FilterContent = r.State.Contents
	r.Exec()
}

func (r *RenderLog) FilterContentAndDraw() {
	r.State.FilterContent = []string{}
	for _, c := range r.State.Contents {
		if strings.Contains(c, r.State.Search) {
			r.State.FilterContent = append(r.State.FilterContent, c)
		}
	}
	r.Exec()
}

func NewRenderLog(app *tview.Application, view *tview.TextView, state *State) *RenderLog {
	return &RenderLog{
		App:   app,
		View:  view,
		State: state,
	}
}
