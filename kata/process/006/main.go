package main

import (
	"fmt"
	"log"
	"os/exec"
)


func main() {
    cmd := exec.Command("ls", "-as")
    result, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(result))
}