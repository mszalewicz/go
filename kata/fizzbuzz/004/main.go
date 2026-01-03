package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	usage := "Usage: program <number>"

	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(usage)
	}

	mapping := []struct {
		denominator int
		text        string
	}{
		{3, "fizz"},
		{5, "buzz"},
		{7, "kezz"},
	}

	firstWord := true
	var sb strings.Builder

	for _, entry := range mapping {
		if input % entry.denominator == 0 {
			if !firstWord {
				sb.WriteString(" ")
			}

			sb.WriteString(entry.text)
			firstWord = false
		}
	}

	println(sb.String())
}
