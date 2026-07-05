package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	command := exec.Command("ls", "-ls")
	output, err := command.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))

}
