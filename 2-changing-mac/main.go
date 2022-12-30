package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	selectedInterface := "eth0"
	runCommand("sudo", "ifconfig", selectedInterface, "down")

}

//Variadic function to accept arguments
func runCommand(cmd string, args ...string) error {

	commandObj := exec.Command(cmd, args...)
	commandObj.Stdout = os.Stdout
	commandObj.Stderr = os.Stderr
	commandObj.Stdin = os.Stdin
	err := commandObj.Run()
	if err != nil {
		log.Fatal(err)

	}
	return nil
}
