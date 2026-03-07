package main

import (
	"fmt"
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

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(usage)
	}

	result := checkNumber(number)

	fmt.Println(result)
}

type Case struct {
	denominator int
	text        string
}

func checkNumber(number int) string {
	var sb strings.Builder

	mapping := []Case{
		{3, "fizz"},
		{5, "buzz"},
		{7, "bring"},
	}

	for i := range mapping {
		if number%mapping[i].denominator == 0 {
			sb.WriteString(mapping[i].text)
			sb.WriteString(" ")
		}
	}

	return sb.String()
}
