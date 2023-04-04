package webtty

import (
	"io"
)

// Slave represents a PTY slave, typically it's a local command.
type Slave interface {
	io.ReadWriter

	GetHistory() []byte
	PushToHistory(data []byte)
	Seek(offset int64) (oldPosition int64, err error)

	// WindowTitleVariables returns any values that can be used to fill out
	// the title of a terminal.
	WindowTitleVariables() map[string]interface{}

	// ResizeTerminal sets a new size of the terminal.
	ResizeTerminal(columns int, rows int) error
}
