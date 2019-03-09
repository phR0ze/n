# term
Package term provides some simple terminal helper functions. Go's built in support is extremely
limited when it comes to working with terminal input. For example go is unable out of the box to
read input data without a return.

### Table of Contents
* [Usage](#usage)
  * [Read from TTY](#read-from-tty)
  * [Listen for SIGWINCH](#listen-for-sigwinch)
* [Research](#research)
  * [Termios](#termios)
    * [Opening a TTY](#opening-a-tty)

## Usage <a name="usage"></a>

### Listen for SIGWINCH <a name="listen-for-sigwinch"></a>
```go
var err error
var tty *term.TTY
if tty, err = term.Open(); err != nil {
	return
}

go func() {
	for {
		select {
		case <-tty.SigWinSizeChan:
			if w, h, err := tty.Size(); err == nil {
				// take some action
			}
		}
	}
}()
```

### Read from TTY <a name="read-from-tty"></a>
```go
var err error
var tty *term.TTY
if tty, err = term.Open(); err != nil {
	return
}
defer tty.Close()

// Read char
c, _ := tty.ReadChar()
fmt.Println(c)
```

## Research <a name="research"></a>
First why would you want to do that?  Well it turns out that Go doesn't have the ability out of
the box to be able to read cli input without first having enter pressed. This is extremely
inconvenient when you just want your app to wait for any key to be pressed, not only the enter
key. To this end I'm putting a little time into understanding how to handle this and writing
a few helper functions for this use case.

* https://github.com/golang/crypto/blob/master/ssh/terminal/util.go
* https://en.wikibooks.org/wiki/Serial_Programming/termios
* https://godoc.org/golang.org/x/sys/unix
* https://github.com/nsf/termbox-go
* https://github.com/mattn/go-tty

### Termios <a name="termios">
Termios is the newer Unix API for terminal I/O. An app that interacts with the terminal typically
will ***open*** a serial device with the standard Unix system call open, configure communication
paramaters, make ***read*** and ***write*** calls then finish with ***close***.

There are two primary modes termios provides:
* ***Cooked/Canonical Mode*** - input is assembled into lines and special characters are processed
* ***Raw/Non-Canonical Mode*** - input is not processed in anyway, just characters are emitted.

Configuration is done via ***termios*** flags

#### Opening a TTY <a name="opening-a-tty">
You have the choice to open the TTY for reading and/or writing.
