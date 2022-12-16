package command

type Command int

const (
	Backspace Command = iota
	CursorDown
	CursorLeft
	CursorRight
	CursorUp
	Delete
	End
	Home
	Insert
	Open
	PageDown
	PageUp
	Quit
	Save
)
