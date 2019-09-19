package main

import (
	"fmt"

	"github.com/phR0ze/n/pkg/sys"
)

func main() {
	user, _ := sys.CurrentUser()
	fmt.Printf("UID: %d\n", user.UID)
	fmt.Printf("GID: %d\n", user.GID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Home: %s\n", user.Home)
	fmt.Printf("IsRoot: %v\n", user.IsRoot())
	fmt.Printf("Real UID: %d\n", user.RealUID)
	fmt.Printf("Real GID: %d\n", user.RealGID)
	fmt.Printf("Real Name: %s\n", user.RealName)
	fmt.Printf("Real Home: %s\n", user.RealHome)
	// sys.WriteString("sudo.txt", "foo")

	// validate that sudo is dropped
	fmt.Printf("\ndrop sudo\n\n")
	sys.DropSudo()
	user, _ = sys.CurrentUser(false)

	fmt.Printf("UID: %d\n", user.UID)
	fmt.Printf("GID: %d\n", user.GID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Home: %s\n", user.Home)
	fmt.Printf("IsRoot: %v\n", user.IsRoot())
	fmt.Printf("Real UID: %d\n", user.RealUID)
	fmt.Printf("Real GID: %d\n", user.RealGID)
	fmt.Printf("Real Name: %s\n", user.RealName)
	fmt.Printf("Real Home: %s\n", user.RealHome)

	// validate external system calls are in the euid perms
	// sys.WriteString("user.txt", "foo")
	// if _, err := sys.ExecOut("echo test > foo"); err != nil {
	// 	fmt.Printf("%+v\n", err)
	// }

	// validate that sudo is raised
	fmt.Printf("\nsudo\n\n")
	if err := sys.Sudo(); err != nil {
		fmt.Printf("%+v\n", err)
	}
	user, _ = sys.CurrentUser(false)

	fmt.Printf("UID: %d\n", user.UID)
	fmt.Printf("GID: %d\n", user.GID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Home: %s\n", user.Home)
	fmt.Printf("IsRoot: %v\n", user.IsRoot())
	fmt.Printf("Real UID: %d\n", user.RealUID)
	fmt.Printf("Real GID: %d\n", user.RealGID)
	fmt.Printf("Real Name: %s\n", user.RealName)
	fmt.Printf("Real Home: %s\n", user.RealHome)
	// sys.WriteString("newsudo.txt", "foo")
}
