package terminal

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	ClearScreen   = "\u001Bc"
	CursorHide    = "\u001B[?25l"
	CursorShow    = "\u001B[?25h"
	CursorTopLeft = "\u001B[H"
	EraseLine     = "\u001B[K"

	SetCursorFmt = "\u001B[%d;%dH"
)

type Terminal struct {
	fd      uintptr
	termios *syscall.Termios
}

type Size struct {
	Rows    uint16
	Columns uint16
	XPixels uint16
	YPixels uint16
}

func New(fd uintptr) (*Terminal, error) {
	var t = &Terminal{
		termios: &syscall.Termios{},
	}

	if _, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		fd,
		syscall.TCGETS,
		uintptr(unsafe.Pointer(t.termios)),
	); err != 0 {
		return nil, fmt.Errorf("ioctl syscall: %w", err)
	}

	return t, nil
}

// Reset the terminal into the original termios.
func (t *Terminal) Reset() error {
	if _, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		t.fd,
		uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(t.termios)),
	); err != 0 {
		return fmt.Errorf("ioctl syscall: %w", err)
	}

	return nil
}

// SetRaw enables raw mode.
func (t *Terminal) SetRaw() error {
	var termios = *t.termios
	termios.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
	termios.Oflag &^= syscall.OPOST
	termios.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
	termios.Cflag &^= syscall.CSIZE | syscall.PARENB
	termios.Cflag |= syscall.CS8
	termios.Cc[syscall.VMIN] = 1
	termios.Cc[syscall.VTIME] = 0

	if _, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		t.fd,
		uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(&termios)),
	); err != 0 {
		return fmt.Errorf("ioctl syscall: %w", err)
	}

	return nil
}

func (t *Terminal) size() (*Size, error) {
	var s = &Size{}
	if _, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		t.fd,
		syscall.TIOCGWINSZ,
		uintptr(unsafe.Pointer(s)),
	); err != 0 {
		return nil, fmt.Errorf("ioctl syscall: %w", err)
	}

	return s, nil
}

// Size returns the zero-indexed size of the terminal.
func (t *Terminal) Size() (*Size, error) {
	size, err := t.size()
	if err != nil {
		return nil, err
	}

	size.Columns--
	size.Rows--

	return size, nil
}
