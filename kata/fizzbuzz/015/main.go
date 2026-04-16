package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	usage := "Usage: programm <number>"

	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Error:", err.Error(), "\n", usage)
	}

	mapping := []struct {
		denominator int
		text        string
	}{
		{denominator: 3, text: "fizz"},
		{denominator: 5, text: "buzz"},
		{denominator: 7, text: "shrug"},
	}

	for _, entry := range mapping {
		if number%entry.denominator == 0 {
			fmt.Println(entry.text)

		}
	}
}
