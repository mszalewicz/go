package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: program <number>")
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Usage: program <number>")
	}

	mapping := []struct {
		denominator int
		message     string
	}{
		{3, "fizz"},
		{5, "buzz"},
		{7, "secret"},
	}

	for _, entry := range mapping {
		if number%entry.denominator == 0 {
			fmt.Println(entry.message)
		}
	}
}
