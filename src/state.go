package src

import "sync"

type FocusMode int

const (
	Search FocusMode = iota
	ViewLog
	Command
)

type State struct {
	Sync          sync.RWMutex
	FocusOn       FocusMode
	Search        string
	SearchEvent   chan string
	Contents      []string
	FilterContent []string
}

func NewState() *State {
	return &State{
		FocusOn:     Command,
		Search:      "",
		SearchEvent: make(chan string, 1),
		Contents:    nil,
		Sync:        sync.RWMutex{},
	}
}
