package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	output, err := exec.Command("ps", "-ax").Output()

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(output), "\n")

	for index, line := range lines {
		if index < 10 {
			fields := strings.Fields(line)

			for _, field := range fields {
				fmt.Printf("%s ", field)
			}
			fmt.Println()
		}
	}

}
