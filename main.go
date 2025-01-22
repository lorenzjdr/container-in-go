package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Checks for user arguements
func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("messed up lol")
	}
}

// Re-run the current program as a child
func parent() {
	// /proc/self/exe is a symlink to the current program
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	// Set the command's standard input, output, and error to the current process's standard input, output, and error
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and check for errors
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

// Run the user's program 
func child() {
	
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}
// Error handling
func must(err error) {
	if err != nil {
		panic(err)
	}
}