package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
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

	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Cloneflags is a bit mask that tells the kernel what to do
		// CLONE_NEWUTS: Create a new UTS namespace
		// CLONE_NEWPID: Create a new PID namespace
		// CLONE_NEWNS: Create a new mount namespace
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

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
	// Mount the root
	must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	// Create a directory for the old root filesystem
	must(os.MkdirAll("rootfs/oldrootfs", 0700))
	// Pivot the root filesystem to the new root
	must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	// Change the current working directory to the new root
	must(os.Chdir("/"))

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