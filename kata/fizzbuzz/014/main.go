package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	usage := "Usage: program <number>"

	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(usage)
	}

	mapping := []struct {
		denominator int
		text        string
	}{
		{denominator: 3, text: "fizz"},
		{denominator: 5, text: "buzz"},
		{denominator: 7, text: "YES"},
	}

	for _, entry := range mapping {
		if number%entry.denominator == 0 {
			fmt.Println(entry.text)
		}
	}
}
