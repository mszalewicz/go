package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	usage := "Usage: program <number>\n"

	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(usage)
	}

	mapping := []struct {
		int
		string
	}{
		{3, "fizz"},
		{5, "buzz"},
		{7, "drink"},
	}

	for _, entry := range mapping {
		if number%entry.int == 0 {
		    fmt.Println(entry.string)
		}
	}
}
