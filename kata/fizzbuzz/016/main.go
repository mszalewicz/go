package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type tuple struct {
	denominator int
	text        string
}

func main() {
	usage := "Usage: application <int>"

	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	numerator, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(usage)
	}

	mapping := []tuple{
		{3, "fizz"},
		{5, "buzz"},
		{7, "low"},
	}

	for _, entry := range mapping {
		if numerator%entry.denominator == 0 {
			fmt.Println(entry.text)
		}
	}
}
