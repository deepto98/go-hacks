package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	selectedInterface := "eth0"

	//Change MAC
	//1. Shut down interface
	runCommand("sudo", "ifconfig", selectedInterface, "down")
	//2. Change MAC
	runCommand("sudo", "ifconfig", selectedInterface, "hw", "ether", "60:45:bd:1e:8d:fe")
	//3. Turn on interface
	runCommand("sudo", "ifconfig", selectedInterface, "up")

}

// Variadic function to accept arguments
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
