package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	command := "date"

	commandObj := exec.Command(command)

	commandObj.Stdout = os.Stdout
	commandObj.Stderr = os.Stderr

	err := commandObj.Run()
	if err != nil {
		log.Fatal(err)

	}

}
