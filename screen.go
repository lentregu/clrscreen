package screen

import (
	"log"
	"os"
	"os/exec"
)

// Enum to describe the operating system to which the clear function must be created
const (
	UNIX    = iota
	LINUX
	DARWIN
	WINDOWS
	OTHER
)

// ClearWindow is a type to define the signature of clear functions
type ClearWindow func()

var commands map[int]string

func init() {
	commands = map[int]string{
		UNIX:    "clear",
		LINUX:   "clear",
		DARWIN:  "clear",
		WINDOWS: "cls",
		OTHER:   "clear",
	}
}

//NewClearScreenFunction is a constructor to create a function to clear screen in the operating system received in the opSys param.
func NewClearScreenFunction(opSys int) ClearWindow {

	if _, ok := commands[opSys]; !ok {
		opSys = OTHER
		log.Printf("WARNING: The OS is not supported, there is not any guarantees that the clean function works fine.")
	}

	return func() {
		cmd := exec.Command(commands[opSys])
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
