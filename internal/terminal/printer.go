package terminal

import (
	"fmt"
	"os"
)

// PrintInfo formats str with args and prints to stdout
func PrintInfo(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, format+"\n", a...)
}

// PrintError formats err and prints to stderr
func PrintError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}
