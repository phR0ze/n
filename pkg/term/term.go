package term

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"unicode"

	"golang.org/x/sys/unix"
)

type readOpts struct {
	echo bool // turn on echoing intput back to output
	mask bool // turn on emitting asterisks instead of echoing output
}

// TTY provides interaction with the system tty
type TTY struct {
	infile  *os.File // input file handle
	outfile *os.File // output file handle

	termios *unix.Termios // copy of termios state

	Stdin          *bufio.Reader  // input reader
	Stdout         *bufio.Writer  // output writer
	SigWinSizeChan chan os.Signal // signal channel for unix.SIGWINCH
}

// AnyKey waits for any key to be pressed before returning.
func AnyKey() (err error) {
	var tty *TTY
	if tty, err = Open(); err != nil {
		return
	}
	defer tty.Close()

	// Read single rune
	tty.ReadRune()

	return
}

// Prompt prints out the given message and waits for any key to be pressed before returning.
func Prompt(msg string) (err error) {
	var tty *TTY
	if tty, err = Open(); err != nil {
		return
	}
	defer tty.Close()

	// Read single rune
	fmt.Println(msg)
	tty.ReadRune()

	return
}

// WaitForKey waits the given key to be pressed
func WaitForKey(key byte) (err error) {
	var tty *TTY
	if tty, err = Open(); err != nil {
		return
	}
	defer tty.Close()

	// Read until given key is pressed
	tty.Stdin.ReadString(key)

	return
}

// Open a new TTY helper instance
// https://github.com/golang/crypto/blob/master/ssh/terminal/util.go
func Open() (tty *TTY, err error) {
	tty = &TTY{}

	// Setup stdin/stdout reader/writer
	if tty.infile, err = os.Open("/dev/tty"); err != nil {
		return
	}
	if tty.outfile, err = os.OpenFile("/dev/tty", unix.O_WRONLY, 0); err != nil {
		return
	}
	tty.Stdin = bufio.NewReader(tty.infile)
	tty.Stdout = bufio.NewWriter(tty.outfile)

	// Save termios current state, errors out if tty.infile is not a valid terminal
	if tty.termios, err = unix.IoctlGetTermios(int(tty.infile.Fd()), ioctlReadTermios); err != nil {
		return
	}

	// Configure termios to turn off terminal extras to make automation simpler
	termios := *tty.termios
	termios.Iflag &^= unix.ISTRIP // turn off stripping parity bits
	termios.Iflag &^= unix.INLCR  // turn off NL to CR conversion
	termios.Iflag &^= unix.IGNCR  // turn off ignoring CR
	termios.Iflag &^= unix.ICRNL  // turn off CR to NL conversion so we get return symbols
	termios.Iflag &^= unix.IXON   // turn off ^S/^Q flow control to avoid a bad state
	termios.Lflag &^= unix.ECHO   // turn off echoing so we can control output for passwords
	termios.Lflag &^= unix.ICANON // turn off canonical mode so that input is made available  immediately
	if err = unix.IoctlSetTermios(int(tty.infile.Fd()), ioctlWriteTermios, &termios); err != nil {
		return
	}

	// Create a channel for receiving SIGWINCH signals from the system
	tty.SigWinSizeChan = make(chan os.Signal, 1)
	signal.Notify(tty.SigWinSizeChan, unix.SIGWINCH)

	return
}

// Size returns the current size of the terminal window.
// Used in conjunction with the SigWinSizeChan one can react to terminal size changes.
func (tty *TTY) Size() (col int, row int, err error) {
	size, err := unix.IoctlGetWinsize(int(tty.infile.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return -1, -1, err
	}
	col, row = int(size.Col), int(size.Row)

	return
}

// Close TTY resources and restore termios save state
func (tty *TTY) Close() (err error) {
	close(tty.SigWinSizeChan)
	err = unix.IoctlSetTermios(int(tty.infile.Fd()), ioctlWriteTermios, tty.termios)
	return
}

// TTY Methods
//--------------------------------------------------------------------------------------------------

// ReadChar from the TTY, blocks until data is present
func (tty *TTY) ReadChar() (string, error) {
	r, _, err := tty.Stdin.ReadRune()
	return string(r), err
}

// ReadRune from the TTY, blocks until data is present
func (tty *TTY) ReadRune() (rune, error) {
	r, _, err := tty.Stdin.ReadRune()
	return r, err
}

// ReadLine reads from the TTY until return is pressed i.e. '\r'
// returned string does not include the trailing '\r'
func (tty *TTY) ReadLine() (result string, err error) {
	result, err = tty.Stdin.ReadString('\r')
	return
}

// ReadString reads from the TTY until return is pressed and echos back to TTY rune by rune
func (tty *TTY) ReadString() (result string, err error) {
	result, err = tty.read(readOpts{echo: true})
	tty.outfile.WriteString("\n")
	return
}

// ReadSensitive reads from TTY until return is pressed, does not echo
func (tty *TTY) ReadSensitive() (result string, err error) {
	result, err = tty.read(readOpts{})
	tty.outfile.WriteString("\n")
	return
}

// ReadPassword reads from TTY until return is pressed, printing out asterisks in place of echo
func (tty *TTY) ReadPassword() (result string, err error) {
	result, err = tty.read(readOpts{echo: true, mask: true})
	tty.outfile.WriteString("\n")
	return
}

// read a string from stdin
func (tty *TTY) read(opts readOpts) (result string, err error) {
	runes := []rune{}

	for {

		// Get rune from TTY
		var x rune
		if x, _, err = tty.Stdin.ReadRune(); err != nil {
			return
		}

		switch x {

		// Done
		case KeyReturn:
			result = string(runes)
			return

		// Handle backspaces/deletes
		case KeyBackSpace, KeyDelete:
			if len(runes) > 0 {
				runes = runes[:len(runes)-1]

				// back up, blot out then back up again
				if opts.echo {
					tty.outfile.WriteString("\b \b")
				}
			}

		// Handle valid characters
		default:

			// Filter out invalid characters
			if unicode.IsPrint(x) {
				runes = append(runes, x)

				// Echo out result if directed
				if opts.echo {
					if opts.mask {
						tty.outfile.WriteString("*")
					} else {
						tty.outfile.WriteString(string(x))
					}
				}
			}
		}
	}
}
