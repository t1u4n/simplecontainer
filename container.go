package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Verify if arguments are valid.
	verifyArgs()

	switch os.Args[1] {
	case "run":
		run()
	}
}

func verifyArgs() {
	if os.Args[1] != "run" {
		panic("Invalid args, see help for usage.")
	}
	fmt.Printf("running %v\n", os.Args[2:])
}

func run() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
