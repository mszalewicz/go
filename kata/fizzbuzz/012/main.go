package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	usage := "Usage: progran <number>"

	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(number)
	}

	mapping := []struct {
		denominator int
		text        string
	}{
		{3, "fizz"},
		{5, "buzz"},
		{7, "test"},
	}

	for _, instance := range mapping {
		if number%instance.denominator == 0 {
			fmt.Println(instance.text)
		}
	}
}
