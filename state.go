package main

type FocusMode int

const (
	Search FocusMode = iota
	ViewLog
	Command
)

type State struct {
	FocusOn FocusMode
	Search  string
}

func NewState() *State {
	return &State{
		FocusOn: Command,
	}
}
