package main

import (
	"log"
	"os"
	"os/exec"
)

//Variadic function to accept arguments
func run(cmd string, args ...string) error {

	commandObj := exec.Command(cmd, args...)
	commandObj.Stdout = os.Stdout
	commandObj.Stderr = os.Stderr

	err := commandObj.Run()
	if err != nil {
		log.Fatal(err)

	}
	return nil
}

func main() {
	cmd := "ls"

	run(cmd, "-l")
}
