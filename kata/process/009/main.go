package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
    cmd := exec.Command("ls", "-ls")

    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(output))

}