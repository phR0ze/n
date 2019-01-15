package main

import (
	"fmt"
	"log"

	"github.com/phR0ze/n/pkg/term"
)

func main() {
	//test_AnyKey()
	//test_WaitForKey()
	//test_ReadRune()
	//test_ReadLine()
	test_ReadString()
	//test_ReadSensitive()
	//test_ReadPassword()
	//test_Size()
}

func test_Size() {
	tty, err := term.Open()
	if err != nil {
		log.Fatal("failed to open TTY")
	}
	defer tty.Close()

	w, h, _ := tty.Size()
	fmt.Printf("Width: %d, Height: %d\n", w, h)
}

func test_ReadPassword() {
	tty, err := term.Open()
	if err != nil {
		log.Fatal("failed to open TTY")
	}
	defer tty.Close()

	for {
		if str, err := tty.ReadPassword(); err != nil {
			log.Fatal("failed to readpassword")
		} else {
			fmt.Println(str)
		}
	}
}

func test_ReadSensitive() {
	tty, err := term.Open()
	if err != nil {
		log.Fatal("failed to open TTY")
	}
	defer tty.Close()

	for {
		if str, err := tty.ReadSensitive(); err != nil {
			log.Fatal("failed to readsensitive")
		} else {
			fmt.Println(str)
		}
	}
}

func test_ReadString() {
	tty, err := term.Open()
	if err != nil {
		log.Fatal("failed to open TTY")
	}
	defer tty.Close()

	for {
		if str, err := tty.ReadString(); err != nil {
			log.Fatal("failed to readstring")
		} else {
			fmt.Println(str)
		}
	}
}

func test_ReadLine() {
	tty, err := term.Open()
	if err != nil {
		log.Fatal("failed to open TTY")
	}
	defer tty.Close()

	for {
		if str, err := tty.ReadLine(); err != nil {
			log.Fatal("failed to readline")
		} else {
			fmt.Println(str)
		}
	}
}

func test_ReadRune() {
	tty, err := term.Open()
	if err != nil {
		log.Fatal("failed to open TTY")
	}
	defer tty.Close()

	for {
		result, err := tty.ReadRune()
		if err != nil {
			log.Fatal("failed to read rune")
		}
		fmt.Println(result)
	}
}

func test_WaitForKey() {
	term.WaitForKey(',')
}

func test_AnyKey() {
	term.AnyKey()
}
