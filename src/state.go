package src

type FocusMode int

const (
	Input FocusMode = iota
	LogTextView
)

type State struct {
	FocusOn FocusMode
}
